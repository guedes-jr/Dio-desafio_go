package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID       int       `json:"id"`
	Nome     string    `json:"nome"`
	Email    string    `json:"email"`
	CriadoEm time.Time `json:"criado_em"`
}

type createClienteRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type updateClienteRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

type patchClienteRequest struct {
	Nome  *string `json:"nome,omitempty"`
	Email *string `json:"email,omitempty"`
}

type memoriaStore struct {
	mu       sync.RWMutex
	seq      int
	clientes map[int]Cliente
	byEmail  map[string]int
}

func newMemoriaStore() *memoriaStore {
	return &memoriaStore{
		clientes: make(map[int]Cliente),
		byEmail:  make(map[string]int),
	}
}

func (s *memoriaStore) list() []Cliente {
	s.mu.RLock()
	defer s.mu.RUnlock()
	res := make([]Cliente, 0, len(s.clientes))
	for _, c := range s.clientes {
		res = append(res, c)
	}
	return res
}

func (s *memoriaStore) get(id int) (Cliente, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c, ok := s.clientes[id]
	return c, ok
}

func (s *memoriaStore) emailExists(email string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.byEmail[strings.ToLower(email)]
	return ok
}

func (s *memoriaStore) create(in createClienteRequest) (Cliente, error) {
	nome := strings.TrimSpace(in.Nome)
	email := strings.TrimSpace(in.Email)
	if nome == "" || email == "" {
		return Cliente{}, errors.New("nome e email são obrigatórios")
	}
	if !strings.Contains(email, "@") {
		return Cliente{}, errors.New("email inválido")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.byEmail[strings.ToLower(email)]; exists {
		return Cliente{}, errors.New("email já cadastrado")
	}

	s.seq++
	now := time.Now().UTC()
	c := Cliente{
		ID:       s.seq,
		Nome:     nome,
		Email:    email,
		CriadoEm: now,
	}
	s.clientes[c.ID] = c
	s.byEmail[strings.ToLower(email)] = c.ID
	return c, nil
}

func (s *memoriaStore) replace(id int, in updateClienteRequest) (Cliente, error) {
	nome := strings.TrimSpace(in.Nome)
	email := strings.TrimSpace(in.Email)
	if nome == "" || email == "" {
		return Cliente{}, errors.New("nome e email são obrigatórios")
	}
	if !strings.Contains(email, "@") {
		return Cliente{}, errors.New("email inválido")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	existing, ok := s.clientes[id]
	if !ok {
		return Cliente{}, os.ErrNotExist
	}

	if strings.ToLower(existing.Email) != strings.ToLower(email) {
		if _, used := s.byEmail[strings.ToLower(email)]; used {
			return Cliente{}, errors.New("email já cadastrado")
		}
		delete(s.byEmail, strings.ToLower(existing.Email))
		s.byEmail[strings.ToLower(email)] = id
	}

	existing.Nome = nome
	existing.Email = email
	s.clientes[id] = existing
	return existing, nil
}

func (s *memoriaStore) patch(id int, in patchClienteRequest) (Cliente, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	existing, ok := s.clientes[id]
	if !ok {
		return Cliente{}, os.ErrNotExist
	}

	if in.Nome != nil {
		n := strings.TrimSpace(*in.Nome)
		if n == "" {
			return Cliente{}, errors.New("nome não pode ser vazio")
		}
		existing.Nome = n
	}

	if in.Email != nil {
		e := strings.TrimSpace(*in.Email)
		if e == "" || !strings.Contains(e, "@") {
			return Cliente{}, errors.New("email inválido")
		}
		old := strings.ToLower(existing.Email)
		newKey := strings.ToLower(e)
		if old != newKey {
			if _, used := s.byEmail[newKey]; used {
				return Cliente{}, errors.New("email já cadastrado")
			}
			delete(s.byEmail, old)
			s.byEmail[newKey] = id
		}
		existing.Email = e
	}

	s.clientes[id] = existing
	return existing, nil
}

func (s *memoriaStore) delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	existing, ok := s.clientes[id]
	if !ok {
		return false
	}
	delete(s.clientes, id)
	delete(s.byEmail, strings.ToLower(existing.Email))
	return true
}

func parseID(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 || id > math.MaxInt32 {
		return 0, errors.New("id inválido")
	}
	return id, nil
}

type api struct {
	store *memoriaStore
}

func (a *api) listClientes(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.store.list())
}

func (a *api) getCliente(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	c, ok := a.store.get(id)
	if !ok {
		writeError(w, http.StatusNotFound, errors.New("cliente não encontrado"))
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (a *api) createCliente(w http.ResponseWriter, r *http.Request) {
	var in createClienteRequest
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	c, err := a.store.create(in)
	if err != nil {
		if err.Error() == "email já cadastrado" {
			writeError(w, http.StatusConflict, err)
			return
		}
		writeError(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Location", "/clientes/"+strconv.Itoa(c.ID))
	writeJSON(w, http.StatusCreated, c)
}

func (a *api) replaceCliente(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	var in updateClienteRequest
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	c, err := a.store.replace(id, in)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writeError(w, http.StatusNotFound, errors.New("cliente não encontrado"))
			return
		}
		if err.Error() == "email já cadastrado" {
			writeError(w, http.StatusConflict, err)
			return
		}
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (a *api) patchCliente(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	var in patchClienteRequest
	if err := decodeJSON(r, &in); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	c, err := a.store.patch(id, in)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			writeError(w, http.StatusNotFound, errors.New("cliente não encontrado"))
			return
		}
		if err.Error() == "email já cadastrado" {
			writeError(w, http.StatusConflict, err)
			return
		}
		writeError(w, http.StatusBadRequest, err)
		return
	}
	writeJSON(w, http.StatusOK, c)
}

func (a *api) deleteCliente(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	ok := a.store.delete(id)
	if !ok {
		writeError(w, http.StatusNotFound, errors.New("cliente não encontrado"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func decodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(dst)
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, err error) {
	type resp struct {
		Erro string `json:"erro"`
	}
	writeJSON(w, code, resp{Erro: err.Error()})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	store := newMemoriaStore()
	api := &api{store: store}

	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(corsMiddleware)

	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}).Methods(http.MethodGet)

	sub := r.PathPrefix("/clientes").Subrouter()
	sub.HandleFunc("", api.listClientes).Methods(http.MethodGet)
	sub.HandleFunc("/", api.listClientes).Methods(http.MethodGet)

	sub.HandleFunc("", api.createCliente).Methods(http.MethodPost)
	sub.HandleFunc("/", api.createCliente).Methods(http.MethodPost)

	sub.HandleFunc("/{id}", api.getCliente).Methods(http.MethodGet)
	sub.HandleFunc("/{id}", api.replaceCliente).Methods(http.MethodPut)
	sub.HandleFunc("/{id}", api.patchCliente).Methods(http.MethodPatch)
	sub.HandleFunc("/{id}", api.deleteCliente).Methods(http.MethodDelete)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("API ouvindo em http://localhost%s", srv.Addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("erro ao subir servidor: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Println("servidor finalizado")
}

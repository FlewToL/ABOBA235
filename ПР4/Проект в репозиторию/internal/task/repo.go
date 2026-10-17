package task

import (
	"errors"
	"sync"
	"time"
)

// Возвращаем эту ошибку, если задача отсутствует.
var ErrNotFound = errors.New("task not found")

type Repo struct {
	mu    sync.RWMutex
	seq    int64
	items  map[int64]*Task
}

func NewRepo() *Repo {
	return &Repo{
		items: make(map[int64]*Task),
	}
}

// Получаем список всех задач.
func (r *Repo) List() []*Task {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]*Task, 0, len(r.items))
	for _, t := range r.items {
		out = append(out, t)
	}

	return out
}

// Ищем задачу по идентификатору.
func (r *Repo) Get(id int64) (*Task, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}

	return t, nil
}

// Создаём задачу с новым идентификатором.
func (r *Repo) Create(title string) *Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.seq++
	now := time.Now()

	t := &Task{
		ID:        r.seq,
		Title:     title,
		CreatedAt: now,
		UpdatedAt: now,
		Done:      false,
	}

	r.items[t.ID] = t
	return t
}

// Обновляем заголовок, статус и время изменения.
func (r *Repo) Update(id int64, title string, done bool) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	t, ok := r.items[id]
	if !ok {
		return nil, ErrNotFound
	}

	t.Title = title
	t.Done = done
	t.UpdatedAt = time.Now()

	return t, nil
}

// Удаляем задачу из хранилища.
func (r *Repo) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.items[id]; !ok {
		return ErrNotFound
	}

	delete(r.items, id)
	return nil
}
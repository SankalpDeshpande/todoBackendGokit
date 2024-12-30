package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // PostgreSQL driver
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type TodoService interface {
	Create(ctx context.Context, todo Todo) (string, error)
	Get(ctx context.Context, id string) (Todo, error)
	Update(ctx context.Context, todo Todo) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]Todo, error)
}

type todoService struct {
	db *sql.DB
}

func NewTodoService(db *sql.DB) TodoService {
	return &todoService{db: db}
}

func (s *todoService) Create(ctx context.Context, todo Todo) (string, error) {
	validStatuses := map[string]bool{
		"pending":     true,
		"in-progress": true,
		"completed":   true,
	}

	if !validStatuses[todo.Status] {
		return "", fmt.Errorf("invalid status: %s", todo.Status)
	}
	todo.ID = uuid.New().String()
	query := `INSERT INTO tasks (id, title, status) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, todo.ID, todo.Title, todo.Status)
	if err != nil {
		log.Printf("Error creating TODO: %v", err)
		return "", err
	}
	return todo.ID, nil
}

func (s *todoService) Get(ctx context.Context, id string) (Todo, error) {
	query := `SELECT id, title, status FROM tasks WHERE id = $1`
	var todo Todo
	err := s.db.QueryRowContext(ctx, query, id).Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, fmt.Errorf("todo not found")
		}
		log.Printf("Error retrieving TODO: %v", err)
		return Todo{}, err
	}
	return todo, nil
}
func (s *todoService) GetAll(ctx context.Context) ([]Todo, error) {
	rows, err := s.db.Query("SELECT id, title, status FROM tasks")
	if err != nil {
		return []Todo{}, err
	}
	defer rows.Close()

	var tasks []Todo
	for rows.Next() {
		var task Todo
		if err := rows.Scan(&task.ID, &task.Title, &task.Status); err != nil {
			return []Todo{}, fmt.Errorf("internal server error")
		}
		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {
		return []Todo{}, fmt.Errorf("no tasks found")
	}

	return tasks, nil
}
func (s *todoService) Update(ctx context.Context, todo Todo) error {
	query := `UPDATE tasks SET title = $1, status = $2 WHERE id = $3`
	result, err := s.db.ExecContext(ctx, query, todo.Title, todo.Status, todo.ID)
	if err != nil {
		log.Printf("Error updating TODO: %v", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id: %s", todo.ID)
	}
	return nil
}

func (s *todoService) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Printf("Error deleting TODO: %v", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no task found with id: %s", id)
	}
	return nil
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rodaine/table"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var tbl = table.New("ID", "Title", "Status", "CreatedAt", "UpdatedAt")

func Mark(status Status, id string) {
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("erro ao abrir aquivo: ", err)
	}

	var Tasks []Todo
	if err := json.Unmarshal(b, &Tasks); err != nil {
		fmt.Println("Erro ao desserializar json: ", err)
	}

	TaskIdx := -1

	for i, todo := range Tasks {
		if todo.ID == id {
			TaskIdx = i
		}
	}

	if TaskIdx == -1 {
		fmt.Println("Tarefa nao encontrada!")
	}

	Tasks[TaskIdx].Status = status

	save(Tasks)
}

func Start() {
	filename := "data.json"

	// os.O_CREATE | os.O_EXCL garante que só criamos se NÃO existir
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o666)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			// Se o arquivo já existe, não faz nada e sai de cabeça erguida
			return
		}
		fmt.Printf("Erro ao tentar inicializar o sistema: %v\n", err)
		return
	}
	defer file.Close()

	// CRÍTICO: Inicializa o arquivo com um array JSON válido `[]`
	// Sem isso, o json.Unmarshal da próxima função vai quebrar com arquivo em branco
	_, _ = file.WriteString("[]")
	fmt.Printf("Arquivo '%s' inicializado com sucesso!\n", filename)
}

func ListFilter(filter Status) {
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("erro ao abrir aquivo: ", err)
	}

	var Tasks []Todo
	if err := json.Unmarshal(b, &Tasks); err != nil {
		fmt.Println("Erro ao desserializar json: ", err)
	}

	for _, todo := range Tasks {
		if todo.Status == filter {
			tbl.AddRow(todo.ID, todo.Title, todo.Status, todo.CreatedAt, todo.UpdatedAt)
		}
	}

	tbl.Print()
}

func List() {
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("erro ao abrir aquivo: ", err)
	}

	var Tasks []Todo
	if err := json.Unmarshal(b, &Tasks); err != nil {
		fmt.Println("Erro ao desserializar json: ", err)
	}

	for _, todo := range Tasks {
		tbl.AddRow(todo.ID, todo.Title, todo.Status, todo.CreatedAt, todo.UpdatedAt)
	}

	tbl.Print()
}

func Add(title string) error {
	id, err := nextID()
	if err != nil {
		return fmt.Errorf("falha ao gerar próximo ID: %w", err)
	}

	newTask := Todo{
		ID:        strconv.Itoa(id),
		Title:     title,
		Status:    StatusTodo,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var Tasks []Todo
	oldFile, err := os.ReadFile("data.json")
	if err == nil {
		if len(oldFile) > 0 {
			if err := json.Unmarshal(oldFile, &Tasks); err != nil {
				return fmt.Errorf("erro ao ler tarefas existentes: %w", err)
			}
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("erro ao abrir arquivo: %w", err)
	}

	Tasks = append(Tasks, newTask)

	jsonDados, err := json.MarshalIndent(Tasks, "", "    ")
	if err != nil {
		return fmt.Errorf("erro ao serializar lista de tarefas: %w", err)
	}

	if err := os.WriteFile("data.json", jsonDados, 0o644); err != nil {
		return fmt.Errorf("erro ao gravar arquivo: %w", err)
	}

	fmt.Println("Tarefa criada com sucesso!")

	return nil
}

func Update(id string, title string) {
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("erro ao abrir aquivo: ", err)
	}

	var Tasks []Todo
	if err := json.Unmarshal(b, &Tasks); err != nil {
		fmt.Println("Erro ao desserializar json: ", err)
	}

	TaskIndex := -1

	for i, task := range Tasks {
		if task.ID == id {
			TaskIndex = i
		}
	}

	if TaskIndex == -1 {
		fmt.Println("Tarefa Nao encontrada")
		return
	}

	Tasks[TaskIndex].Title = title

	json, err := json.MarshalIndent(Tasks, "", "    ")
	if err != nil {
		fmt.Println("Erro convertendo em JSON: ", err)
	}

	if err := os.WriteFile("data.json", json, 0o644); err != nil {
		fmt.Println("Erro ao escrever no Arquino")
	}
}

func Delete(id string) {
	if !exists(id) {
		fmt.Println("a Tarefa nao existe")
		return
	}
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("erro ao abrir aquivo: ", err)
	}

	var Tasks []Todo
	if err := json.Unmarshal(b, &Tasks); err != nil {
		fmt.Println("Erro ao desserializar json: ", err)
	}

	var newTodoList []Todo

	for _, todo := range Tasks {
		if todo.ID == id {
			continue
		}

		newTodoList = append(newTodoList, todo)
	}

	err = save(newTodoList)
	if err != nil {
		fmt.Println("Erro ao Salvar: ", err)
		return
	}

	fmt.Println("Arquivo deletado com sucesso!")
}

func exists(id string) bool {
	b, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("erro ao abrir aquivo: ", err)
	}

	var Tasks []Todo
	err = json.Unmarshal(b, &Tasks)
	if err != nil {
		return false
	}

	for _, todo := range Tasks {
		if todo.ID == id {
			return true
		}
	}

	return false
}

func save(data []Todo) error {
	json, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Erro convertendo em JSON: ", err)
		return err
	}

	if err := os.WriteFile("data.json", json, 0o644); err != nil {
		fmt.Println("Erro ao escrever no Arquino")
		return err
	}

	return nil
}

func nextID() (int, error) {
	b, err := os.ReadFile("data.json")
	if err != nil {
		// Se o arquivo não existir ainda, começamos do ID 1
		if errors.Is(err, os.ErrNotExist) {
			return 1, nil
		}
		return 0, err
	}

	// Se o arquivo estiver completamente vazio (0 bytes), retorna ID 1 de forma segura
	if len(b) == 0 {
		return 1, nil
	}

	var Tasks []Todo
	err = json.Unmarshal(b, &Tasks)
	if err != nil {
		return 0, err
	}

	// Se a lista estiver vazia, o próximo ID é 1
	if len(Tasks) == 0 {
		return 1, nil
	}

	greaterID := 0
	for _, todo := range Tasks {
		id, _ := strconv.Atoi(todo.ID)
		if id > greaterID {
			greaterID = id
		}
	}

	// Retorna o maior ID encontrado + 1
	return greaterID + 1, nil
}

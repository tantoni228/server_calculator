package conection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	Id      int
	Example string
	Result  int
	Error   string
	Steps   []string
	Status  int
}

type Operation struct {
	Id      int
	Id_task int
	Sign    string
	Num1    int
	Num2    int
	Count   int
	// Это индекс вопроса
}

type Steps struct {
	Id   int
	Num  int
	Step string
}

const (
	username = "root"
	password = "root"
	hostname = "localhost:3306"
	dbname   = "golang"
)

func (s *Operation) Calculate2() (int, error) {
	time.Sleep(1000)
	switch s.Sign {
	case "+":
		return s.Num1 + s.Num2, nil
	case "-":
		return s.Num1 - s.Num2, nil
	case "/":
		if s.Num2 == 0 {
			return 0, fmt.Errorf("На ноль делить нельзя")
		} else {
			return s.Num1 / s.Num2, nil
		}
	case "*":
		return s.Num1 * s.Num2, nil
	default:
		return 0, fmt.Errorf("Недопустимый формат")
	}
}

func CreateDatabase() error {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+hostname+")/"+dbname)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	// Создание таблицы Task
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
        Id INT AUTO_INCREMENT PRIMARY KEY,
        Example VARCHAR(255),
        Result INT,
        Error VARCHAR(255),
        Status INT
    )`)
	if err != nil {
		return err
	}

	// Создание таблицы Operation
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Operation (
        Id INT AUTO_INCREMENT PRIMARY KEY,
		Id_task INT,
        Sign VARCHAR(1),
        Num1 INT,
        Num2 INT,
        Count INT
    )`)
	if err != nil {
		return err
	}

	return nil
}

func WriteTask(task Task) (int, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO tasks (example, status) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(task.Example, task.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func ReadTask(id int) (Task, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		return Task{}, err
	}

	defer db.Close()
	var task Task
	err = db.QueryRow("SELECT id, example FROM tasks WHERE id = ?", id).Scan(&task.Id, &task.Example)
	if err != nil {
		return task, err
	}
	return task, nil
}

func WriteOperation(operation Operation) (int, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO operation (id_task, sign, num1, num2, count) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(operation.Id_task, operation.Sign, operation.Num1, operation.Num2, operation.Count)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func ReadOperation(id int) (Operation, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		return Operation{}, err
	}

	defer db.Close()
	var operation Operation
	err = db.QueryRow("SELECT id, sign, num1, num2 FROM operation WHERE id = ?", id).Scan(&operation.Id, &operation.Sign, &operation.Num1, &operation.Num2)
	if err != nil {
		return operation, err
	}
	return operation, nil
}

func GetOperationsByTaskID(id_task int) ([]Operation, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, id_task, sign, num1, num2, count FROM operation WHERE id_task = ?", id_task)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var operations []Operation
	for rows.Next() {
		var op Operation
		if err := rows.Scan(&op.Id, &op.Id_task, &op.Sign, &op.Num1, &op.Num2, &op.Count); err != nil {
			return nil, err
		}
		operations = append(operations, op)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}

package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/tantoni228/server_calculator/db/conection"
	"github.com/tantoni228/server_calculator/pkg/example"
)

func New(ctx context.Context,
) (http.Handler, error) {
	serveMux := http.NewServeMux()

	

	serveMux.HandleFunc("/calculate", Handler_tasks)

	return serveMux, nil
}

func HandlerHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// функция обработчик, принимает и возвращает значение
func Handler_tasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Обрабатывается запрос на сервере на порту 8080")
	ex := r.URL.Query().Get("example")
	fmt.Fprintf(w, "%s=", ex)
	task := conection.Task{
		Example:   ex,
		Status:  0,
	}

	id_task, err := conection.WriteTask(task)
	if err != nil {
		fmt.Fprint(w, err)
		fmt.Printf("Проблемы с записью task %s\n", err)
	}
	data, err := example.Split(ex, id_task)
	if err != nil {
		fmt.Fprint(w, err)
	}
	fmt.Fprintf(w, "%d\n", data)
	fmt.Fprintf(w, "id_task=%d", id_task)
}


func Handler_steps(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: %s", err)
	}
	task, err := conection.ReadTask(id)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else {
		fmt.Fprintf(w, "%s = \n", task.Example)
	}
	data, err := conection.GetOperationsByTaskID(id)
	if err != nil {
		fmt.Fprintf(w, "Ошибка: %s\n", err)
	}
	for _, val := range data {
		f, err := val.Calculate2() 
		if err != nil {
			fmt.Fprintf(w, "%s\n", err)
		} else {

		}
		fmt.Fprintf(w, "%d %s %d = %d\n", val.Num1, val.Sign, val.Num2, f)
	}
	
	
}

func Handler_operation(w http.ResponseWriter, r *http.Request) {
	id1 := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id1)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	operation, err := conection.ReadOperation(id)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	num, err := operation.Calculate2()
	time.Sleep(5 * time.Second)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	fmt.Fprintf(w, "%d", num)
}

// func Handler_action(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, )
// }



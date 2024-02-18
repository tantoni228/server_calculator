package example

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/tantoni228/server_calculator/db/conection"
)

var servers = map[int]bool{
	1: true,
	2: true,
}

func Server() int {
	for i, _ := range servers {
		if servers[i] {
			return i
		}
	}
    fmt.Println("Ожидание свободного агента")
    time.Sleep(5 * time.Second)
    return Server()
}



var wg sync.WaitGroup 

type Operation struct {
    Id   int
	Sign string
	Num1 int
	Num2 int
	Count int 
	// Это индекс вопроса
}

func Request(port int, id int) (string, error){
    fmt.Printf("Запрос выполнен по запросу http://127.0.0.1:%d/?id=%d\n", port, id)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://127.0.0.1:%d/?id=%d", port, id), nil)
    if err != nil {
        return "", nil
    }

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err.Error())
        return "", nil
    }
	// читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
	    fmt.Println(err)
	    return "", nil
	}

	// выводим тело ответа на экран
	return string(body), nil
}

func Calculate(ch chan []int, id int, count int) {
    for i, val := range servers {
        fmt.Printf("Состояние агента по порту :%d -  %t\n", 8080 + i, val)
    }
    defer wg.Done()
    b := Server()
    // if a == 0 {
    //     time.Sleep(5 * time.Second)
    //     fmt.Println("Ожидание свободного агента")
    //     Calculate(ch, id, count)
    //     return 
    // }
    servers[b] = false
    fmt.Printf("Делаю запрос агенту по порту %d\n", b+8080)
    data, err := Request(b + 8080, id)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Printf("Получил от агента по порту %d: %s\n", b + 8080, data)
    a, err := strconv.Atoi(data)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(a)
    servers[b] = true

    ch <- []int{a, count}
}

func (s *Operation) Calculate2() (int, error){
	switch s.Sign {
	case "+":
		return s.Num1 + s.Num2, nil
	case "-":
		return s.Num1 - s.Num2, nil
	case "/" :
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

func Find_server() {
    
}

func remove(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

func Split(s string, id_task int) (int, error){
    fmt.Println("Начал вычисление")
    lines := ParseEquation(s)
    count := 0
    // коэфицент для подсчета расположения ?
    h:= 0
    index := make([]int, 0) 
    // для хранения индексов с знками операциями * или /
    for i, val := range lines {
        if val == "*" || val == "/" {
            h++
            index = append(index, i)
        }
    }
    if h > 0 {
        j := 0 
        
        // для подсчета смещения
        data := make(chan []int)
        for _, val := range index {
            if lines[val + 1 - j] != "?" && lines[val - 1 - j] != "?" {
                wg.Add(1)
                a, err := strconv.Atoi(lines[val - 1 - j])
                if err != nil {
                    return 0, nil
                }
                b, err := strconv.Atoi(lines[val + 1 - j])
                if err != nil {
                    return 0, nil
                }
                sign := lines[val - j]
                // fmt.Println(a, sign, b)
                lines = remove(lines, val - 1 - j)
                lines = remove(lines, val - j)
                lines[val - 1 - j] =  "?"
                
                obj := conection.Operation{Id_task: id_task, Sign: sign, Num1: a, Num2: b, Count: count}
                fmt.Printf("Запись operation\n")
                fmt.Println(obj)
                id, err := conection.WriteOperation(obj)
                if err != nil {
                    fmt.Printf("Проблемы с записью operation %s\n", err)
                }
                fmt.Printf("Id operation: %d\n", id)
                fmt.Println("Начинаю делать запрос агентам")
                go Calculate(data, id, count)
                j += 2
                count++
                // fmt.Println(lines)
            }
        }
        go func() {
            wg.Wait()
            close(data)
        }()
        h = 0
        for result := range data {
			if result[1] < 0 {
                fmt.Println("Ошибка")
				if result[1] == -1 {
					return 0, errors.New("Деление на ноль невозможно.")
				}
				if result[1] == -2 {
					return 0, errors.New("Недопустимая операция.")
				}
 			}
            i := result[1]
            res := result[0]
            j = 0
            
            for inx, val := range lines {
                if val == "?" {
                    if j + h == i {
                        lines[inx] = strconv.Itoa(res)
                        h++
                    }
                    j++
                }
            }
            // fmt.Println(lines)
        }
    } else {
        if len(lines) == 0 {
            return 0, errors.New("Недопустимое значение")
        }
        a, err := strconv.Atoi(lines[0])
        if err != nil {
            return 0, nil
        }
        if len(lines) == 1{
            return a, nil
        }
        b, err := strconv.Atoi(lines[2])
        if err != nil {
            return 0, nil
        }
        sign := lines[1]
        // fmt.Println(a, sign, b)
        lines = remove(lines, 0)
        lines = remove(lines, 1)
        data := make(chan []int)
        wg.Add(1)
        obj := conection.Operation{Id_task: id_task, Sign: sign, Num1: a, Num2: b, Count: count}
        fmt.Printf("Запись operation\n")
        fmt.Println(obj)
        id, err := conection.WriteOperation(obj)
        if err != nil {
            fmt.Printf("Проблемы с записью operation %s\n", err)
        }
        fmt.Printf("Id operation: %d\n", id)
        fmt.Println("Начинаю делать запрос агентам")
        go Calculate(data, id, count)
        
        go func() {
            wg.Wait()
            close(data)
        }()
        for result := range data {
			if result[1] < 0 {
                fmt.Println("Ошибка")
				if result[1] == -1 {
					return 0, errors.New("Деление на ноль невозможно.")
				}
				if result[1] == -2 {
					return 0, errors.New("Недопустимая операция.")
				}
 			}
            lines[0] = strconv.Itoa(result[0])
        }
    }
    if len(lines) >= 3 {
        return Split(strings.Join(lines, ""), id_task)
    } else {
        str, err := strconv.Atoi(strings.Join(lines, ""))
        if err != nil {
            return 0, nil
        }
        return str, nil
    }
}

func ParseEquation(equation string) []string {
	re := regexp.MustCompile(`\d+|[-+*/]`)
	elements := re.FindAllString(equation, -1)
	return elements
}

package handler

import (
	"errors"
	"fmt"
	"log" 
	"sync"
	"encoding/json"	
	"github.com/in-rajendra-pawar/demysttools/net"
	"github.com/in-rajendra-pawar/demysttools/config"
)

const TODO_OUTPUT_FORMAT = "Title: %s | Completed: %s"

type ToDoGetResponse struct {
  UserId int `json:"userId"`
  ID int `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

func ToDoHandler() error { 
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err.Error())
		return err 
	}

	toDoID := cfg.TodoStartID

	wg := sync.WaitGroup{}
	ch := make(chan int, cfg.MaxGoroutine)
	defer close(ch)
	for i := 0; i < cfg.MaxTodo; i ++ { 
		if ( cfg.TodoIDType == "even" && toDoID % 2 != 0 ) || ( cfg.TodoIDType == "odd" && toDoID % 2 == 0 ) {
				toDoID++
		}
		wg.Add(1)
		ch <- 1
		go GetAndPrintToDo(ch, &wg, toDoID, cfg.Url)
		toDoID++
	}
	wg.Wait()

	return nil
}

func GetAndPrintToDo(c chan int, w *sync.WaitGroup, id int, url string) {
		defer w.Done()
		defer func(){
			<-c
		}()

		toDoGetResponse, err := getToDo(url, id)
		
		if err != nil {
			log.Println(fmt.Sprintf("error occured for Todo ID %v", id))
			return 
		}
		completed := "No"
		if toDoGetResponse.Completed {
			completed = "Yes"
		}  
		output := fmt.Sprintf(TODO_OUTPUT_FORMAT, toDoGetResponse.Title, completed)
		fmt.Println(output, " | ", toDoGetResponse.ID )	
}

func getToDo(url string, todoID int) (*ToDoGetResponse, error) {
	toDoGetResponse := ToDoGetResponse{} 
	
	if url == "" {
		err := errors.New("invalid url")
        log.Fatal(err)
        return &toDoGetResponse, err
    }

	response, err := net.Get(fmt.Sprintf("%s%d", url, todoID))

	if err != nil {
        log.Fatal(err)
        return &toDoGetResponse, err
    }

	
	err = json.Unmarshal(response, &toDoGetResponse)
	if err != nil {
        log.Fatal(err)
        return &toDoGetResponse, err
    }
	
	return &toDoGetResponse, nil
}
package handler

import (
    "testing"
    "fmt"
    "net/http"
    "net/http/httptest"
	  "reflect" 
)
 

// Define a test function
func TestGetToDo(t *testing.T) {
  // Use a table-driven approach to test different scenarios
  tests := []struct {
    name    string   // The name of the test case
    url     string   // The url to pass to the getToDo function
    todoID  int      // The todo id to pass to the getToDo function
    want    *ToDoGetResponse // The expected return value
    expectedError bool     // Whether an error is expected or not
  }{
    {
      name:    "valid url and todo id",
      url:     "https://jsonplaceholder.typicode.com/todos/",
      todoID:  1,
      want:    &ToDoGetResponse{UserId: 1, ID: 1, Title: "delectus aut autem", Completed: false},
      expectedError: false,
    }, 
  }

  // Iterate over the test cases
  for _, test := range tests {
    // Use t.Run to run each subtest
    t.Run(test.name, func(t *testing.T) {
      // Create a mock server to handle the net.Get requests
      server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Write a dummy response based on the todo id
        w.Write([]byte(fmt.Sprintf(`{"userId":1,"id":%s,"title":"delectus aut autem","completed":false}`, r.URL.Path[1:])))
      }))
      // Close the server after the test
      defer server.Close()
      // Replace the url with the mock server url
      test.url = server.URL + "/"
      // Call the getToDo function and check for errors
      got, err := getToDo(test.url, test.todoID)
      
      fmt.Printf("\n %+v, \n Type:%T \n err: %T", got, got, err)
      if (got == nil && err != nil) != test.expectedError {
        t.Errorf("getToDo() error = %v, expectedError %v", got, test.expectedError)
      }
      // Compare the return value with the expected value
      if !reflect.DeepEqual(got, test.want) {
        t.Errorf("getToDo() = %v, want %v", got, test.want)
      }
    })
  }
}

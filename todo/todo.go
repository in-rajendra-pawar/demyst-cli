package todo
import(
	"github.com/in-rajendra-pawar/demysttools/todo"
	"fmt"
)

func GetTodoCommand() Cmd {
	cmd := Cmd{
			"ToDo",
			"call rest api and fetch todo",
			`use
				net todo 10
				
				net call api and print output on terminal
				todo fetch todo
				10 number of todo to fetch and print`,
	}
	return cmd
}
func Execute() error {
	fmt.Println("net todo Execute()")
	return nil
}
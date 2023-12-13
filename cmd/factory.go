package cmd
import(
	"github.com/in-rajendra-pawar/demysttools/todo"
	"errors"
	"strings"
)
func GetCommand(cmdName string, cmdOption string) (CliCommand, error){
	switch {
		case strings.ToLower(cmdName) == "net" && strings.ToLower(cmdOption) == "todo":
			return todo.GetTodoCommand(), nil
		default:
			return nil, errors.New("command not found")
	}
}
// 
package todos

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/store/todo"
	"github.com/valentinRog/sba-todo/ui/utils"
)

type Todos struct{}

var (
	Ul     = utils.AddClass(id, h.Ul)
	Div    = utils.AddClass(id, h.Div)
	Li     = utils.AddClass(id, h.Li)
	Form   = utils.AddClass(id, h.Form)
	Input  = utils.AddClass(id, h.Input)
	Button = utils.AddClass(id, h.Button)
)

func addTodoForm() g.Node {
	return Form(
		g.Attr("hx-post", "/add-todo"),
		g.Attr("hx-target", "#todo-list"),
		g.Attr("hx-swap", "outerHTML"),
		g.Attr("hx-on", "htmx:afterRequest: this.reset()"),
		Input(h.Type("text"), h.Name("text")),
		Button(h.Type("submit"), g.Text("add todo")),
	)
}

func (t Todos) TodoList(todos []todo.Todo) g.Node {
	return Ul(h.ID("todo-list"),
		g.Group(g.Map(todos, func(todo todo.Todo) g.Node {
			return Li(g.Text(todo.Name),
				Button(
					g.Attr("hx-post", fmt.Sprintf("/delete-todo/%d", todo.ID)),
					g.Attr("hx-target", "#todo-list"),
					g.Attr("hx-swap", "outerHTML"),
					g.Text("delete"),
				))
		})))
}

func (t Todos) Todos(todos []todo.Todo) g.Node {
	return Div(t.TodoList(todos), addTodoForm())
}

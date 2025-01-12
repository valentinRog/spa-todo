package signup

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	"github.com/valentinRog/sba-todo/ui/utils"

	_ "embed"
)

var (
	h1     = utils.AddClass(Id, h.H1)
	div    = utils.AddClass(Id, h.Div)
	form   = utils.AddClass(Id, h.Form)
	input  = utils.AddClass(Id, h.Input)
	button = utils.AddClass(Id, h.Button)
	a      = utils.AddClass(Id, h.A)
)

func SignupForm() g.Node {
	return div(
		h.ID("signup-form"),
		form(
			g.Attr("action", "/auth/signup"),
			g.Attr("method", "POST"),
			input(h.Type("text"), h.Name("username")),
			input(h.Type("text"), h.Name("password")),
			button(h.Type("submit"), g.Text("signup")),
		),
		a(
			g.Text("signin"),
			g.Attr("hx-get", "/login/signin-form"),
			g.Attr("hx-target", "#signup-form"),
			g.Attr("hx-swap", "outerHTML"),
		),
	)
}
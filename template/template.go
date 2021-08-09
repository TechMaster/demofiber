package template

func Init() *Engine {

	tmpl := New("./views", ".html")

	// Reload the templates on each render, good for development
	tmpl.Reload(true) // Optional. Default: false

	// Debug will print each template that is parsed, good for debugging
	tmpl.Debug(true) // Optional. Default: false

	tmpl.AddFunc("foo", Foo)

	tmpl.SetDefaultTemplate("layout/layout")

	return tmpl
}

func Foo() string {
	return "<h1>Foo</h1>"
}

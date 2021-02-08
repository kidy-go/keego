package keego

type Application struct {
	Config *Config
	Router Router
}

func New() *Application {
	return &Application{}
}

func (app *Application) Bootstrap() *Application {
	app.Config = LoadEnv(".")
	app.Router = NewRouter(app)
	return app
}

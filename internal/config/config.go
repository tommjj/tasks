package config

type (
	App struct {
		TASK_FILE string
	}
)

func New() App {
	return App{
		TASK_FILE: "./tasks.json",
	}
}

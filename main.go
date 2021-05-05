package main

import app2 "dig-playground/app"

func main() {
	app, err := app2.NewApp()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	{% if cookiecutter.use_cobra_cmd == "n" %}"flag"
	"fmt"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/version"{% endif %}
	{% if cookiecutter.use_cobra_cmd == "y" %}"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/cmd"{% endif %}
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const name = "echoserver"

func main() {

    {% if cookiecutter.use_cobra_cmd == "y" %}
    cmd.Execute()
	{% else %}
	versionFlag := flag.Bool("version", false, "Version")
	flag.Parse()

	if *versionFlag {
		fmt.Println("Build Date:", version.BuildDate)
        fmt.Println("Git Commit:", version.GitCommit)
        fmt.Println("Version:", version.Version)
        fmt.Println("Go Version:", version.GoVersion)
        fmt.Println("OS / Arch:", version.OsArch)
		return
	}
	fmt.Println("Hello.")
    {% endif %}

	l := log.New(os.Stdout, "", 0)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)
	app := NewApp(os.Stdin, l)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}
}

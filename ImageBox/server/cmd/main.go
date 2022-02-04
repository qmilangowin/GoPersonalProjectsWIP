package main

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/qmilangowin/imagebox/pkg/config"
	logging "github.com/qmilangowin/imagebox/pkg/logging"
	"github.com/qmilangowin/imagebox/pkg/web"
)

// ServerApplication will make objects available to our handlers via
// dependency injection
type ServerApplication struct {
	Log           *logging.Logger
	templateCache map[string]*template.Template
}

func main() {

	var textConfig config.TextFormatConfig = config.NewTextConfig()
	var logConfig config.LoggingConfig = config.NewLoggingConfig()
	addr := flag.String("port", "8080", "Sets HTTP Port number")
	logDebug := flag.Bool("debug-log", true, "Sets logging to debug (on by default)")
	logOutput := flag.Bool("log-console", true, "Sets output for logging. Default: os.Stdout, set to false to dispose")

	flag.Parse()
	//initialize logging and variables
	*addr = ":" + *addr
	var w io.Writer
	//TODO: give option to write to bytes.Buffer if needed
	//var buf bytes.Buffer
	if *logOutput {
		w = os.Stdout
	} else {
		w = ioutil.Discard
	}

	appLog := logConfig.SetLogging(w, *logDebug)

	//Initialize template cache for UI elements
	//TODO: replace with Hugo or other front-end later
	templateCache, err := web.NewTemplateCache("./ui/html/")
	if err != nil {
		appLog.Log.PrintError(err)
	}

	app := &ServerApplication{
		Log:           appLog.Log,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Handler:      app.routes(),
		Addr:         *addr,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	//start the server
	go func() {
		textFormat := textConfig.SetTextFormatting("cyan")
		fmt.Println(textFormat.Color, "Starting server on port", *addr, textFormat.Reset)
		appLog.Log.Infoprintf("Starting server on: %s", *addr)
		if err := srv.ListenAndServe(); err != nil {
			appLog.Log.Infoprintf("Shutting down server %s", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	//signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	appLog.Log.Infoprintf("Received Terminate, graceful shutdown %s", sig)
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(tc)
}

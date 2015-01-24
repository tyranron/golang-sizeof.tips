package app

import (
	"net/http"
	"os"
	"runtime"
	"sync"
	"syscall"

	"github.com/gophergala/golang-sizeof.tips/internal/log"

	daemon "github.com/tyranron/daemonigo"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Run() (exitCode int) {
	switch isDaemon, err := daemon.Daemonize(); {
	case !isDaemon:
		return
	case err != nil:
		log.StdErr("could not start daemon, reason -> %s", err.Error())
		return 1
	}
	waiter := &sync.WaitGroup{}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Gala!"))
	})
	waiter.Add(1)
	go func() {
		defer waiter.Done()
		http.ListenAndServe(":7777", nil)
	}()

	notifyParentProcess()

	waiter.Wait()
	return
}

// Notifies parent process that everything is OK.
func notifyParentProcess() {
	syscall.Kill(os.Getppid(), syscall.SIGUSR1) // todo: error checking
}

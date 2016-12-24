package yesman

import (
	"flag"
	"fmt"
	"log"
	"log/syslog"
	"os"
	"os/signal"
	"syscall"
)

// Run is main function of yesman.
// Output args's string to stdin.
func Run(ver string, rev string) int {
	// setup log setting
	w, err := syslog.New(syslog.LOG_INFO, "yesman")
	if err != nil {
		panic(err)
	}
	log.SetOutput(w)

	// Trap signal
	signals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}
	go trapSignal(signals)

	log.Printf("[begin] yesman.Run:\n%v\n", os.Args)
	// Settings for CLI options
	count := flag.Int("count", 0, "Limit to output time.")
	version := flag.Bool("version", false, "Show version information.")
	flag.Parse()

	// version options
	if *version == true {
		fmt.Printf("version: %s, revision: %s", ver, rev)
		return 0
	}

	// output args to Stdout
	if *count == 0 {
		for {
			fmt.Printf("%s\n", flag.Args()[0])
		}
	} else {
		for index := 0; index < *count; index++ {
			fmt.Printf("%s\n", flag.Args()[0])
		}
	}

	// Success
	log.Printf("[end  ] yesman.Run:\nExit code: 0\n")
	return 0
}

func trapSignal(signals []os.Signal) {
	log.Printf("[begin] yesman.trapSignal:\n%v\n", signals)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, signals...)
	sig := <-sigCh
	log.Printf("[end  ] yesman.trapSignal:\nReceive syscall signal: %v\n", sig)
	os.Exit(1)
}

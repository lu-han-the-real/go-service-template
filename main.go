package main

import (
	"context"
	"flag"
	"fmt"
	"go-service-template/server"
	"go-service-template/user"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/go-kit/kit/log/logrus"
)

func main() {
	fs := flag.NewFlagSet("addcli", flag.ExitOnError)
	var (
		httpAddr = fs.String("http-addr", ":8080", "HTTP address of addsvc")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags] <a> <b>")
	fs.Parse(os.Args[1:])
	if len(fs.Args()) == 1 {
		fs.Usage()
		os.Exit(1)
	}

	if *httpAddr == "" {
		os.Exit(1)
	}

	store := user.NewStore()
	userServer := user.NewServer(logrus.Logger{}, store)

	httpServer := server.NewHTTPServer(context.Background(), server.MakeEndpoints(userServer))
	err := http.ListenAndServe(*httpAddr, httpServer)
	if err != nil {
		fmt.Println(err)
	}
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}

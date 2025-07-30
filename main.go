package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/colesnodgrass/abcdk/cmds"
	"io"
	"os"
	"os/signal"
	"syscall"
)

type Cmd struct {
	Spec     cmds.SpecCmd     `cmd:""`
	Check    cmds.CheckCmd    `cmd:""`
	Discover cmds.DiscoverCmd `cmd:""`
	Read     cmds.ReadCmd     `cmd:""`
	Write    cmds.WriteCmd    `cmd:""`
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

}

func run(ctx context.Context) error {
	var root Cmd
	parser, err := kong.New(
		&root,
		kong.Name("abcdk"),
		kong.BindToProvider(bindCtx(ctx)),
		kong.BindToProvider(bindWriter(WriterLn{w: os.Stdout})),
		kong.BindToProvider(bindReader(os.Stdin)),
	)
	if err != nil {
		return err
	}
	parsed, err := parser.Parse(os.Args[1:])
	if err != nil {
		return err
	}

	return parsed.Run()
}

func bindCtx(ctx context.Context) func() (context.Context, error) {
	return func() (context.Context, error) {
		return ctx, nil
	}
}

func bindWriter(w io.Writer) func() (io.Writer, error) {
	return func() (io.Writer, error) {
		return w, nil
	}
}

func bindReader(r io.Reader) func() (io.Reader, error) {
	return func() (io.Reader, error) {
		return r, nil
	}
}

// WriterLn wraps a writer to always write a newline character
type WriterLn struct {
	w io.Writer
}

var ln = []byte("\n")

func (w WriterLn) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	if err != nil {
		return n, err
	}

	if !bytes.HasSuffix(p, ln) {
		m, err := w.w.Write(ln)
		if err != nil {
			return n, err
		}
		n += m
	}

	return n, nil
}

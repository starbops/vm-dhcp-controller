package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

const defaultPort = 8080

const tmpDir = "/tmp/vm-dhcp-controller"

func NewServer(ctx context.Context) error {
	if err := createTmpDir(); err != nil {
		return err
	}
	return newServer(ctx, tmpDir)
}

func newServer(ctx context.Context, path string) error {
	defer os.RemoveAll(tmpDir)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", defaultPort),
		Handler: http.FileServer(http.Dir(path)),
	}

	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		return srv.ListenAndServe()
	})

	eg.Go(func() error {
		<-ctx.Done()
		return srv.Shutdown(ctx)
	})

	return eg.Wait()
}

func createTmpDir() error {
	if _, err := os.Stat(tmpDir); os.IsNotExist(err) {
		return os.Mkdir(tmpDir, 0755)
	} else {
		return err
	}
}

// Address returns the address for vm-dhcp url. For local testing set env variable
// SVC_ADDRESS to point to a local endpoint
func Address() string {
	address := "harvester-vm-dhcp-controller.harvester-system.svc"
	if val := os.Getenv("SVC_ADDRESS"); val != "" {
		address = val
	}
	return address
}

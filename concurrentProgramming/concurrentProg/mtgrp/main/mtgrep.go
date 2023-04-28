package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"mgrep/worker"
	"mgrep/worklist"
	"os"
	"path/filepath"
	"sync"
)

func discoverDirs(wl *worklist.WorkList, path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			nextPath := filepath.Join(path, entry.Name())
			discoverDirs(wl, nextPath)
		} else {
			wl.Add(worklist.NewJob(filepath.Join(path, entry.Name())))
		}
	}
}

var args struct {
	SearchTerm string `arg:positional,required`
	SearchDir  string `arg:positional`
}

func main() {
	arg.MustParse(&args)
	var workerWg sync.WaitGroup
	wl := worklist.New(100)
	results := make(chan worker.Result, 100)
	numWorkers := 10
	workerWg.Add(1)

	go func() {
		defer workerWg.Done()
		discoverDirs(&wl, args.SearchDir)
		wl.Finalize(numWorkers)
	}()

	for i := 0; i < numWorkers; i++ {
		workerWg.Add(1)
		go func() {
			defer workerWg.Done()
			for {
				workEntry := wl.Next()
				if workEntry.Path != "" {
					workerResult := worker.FindInFile(workEntry.Path, args.SearchTerm)
					if workerResult != nil {
						for _, r := range workerResult.Inner {
							results <- r
						}
					}
				} else {
					return
				}
			}
		}()
	}

	blockWorkersWg := make(chan struct{})

	go func() {
		workerWg.Wait()
		close(blockWorkersWg)
	}()

	var displayWg sync.WaitGroup

	displayWg.Add(1)

	go func() {
		for {
			select {
			case r := <-results:
				fmt.Printf("%v[%v]:%v", r.Path, r.LineNum, r.Line)
			case <-blockWorkersWg:
				if len(results) == 0 {
					displayWg.Done()
					return
				}
			}
		}
	}()

	displayWg.Wait()
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	basePath        string
	envFilePath     string
	exampleFilePath string

	watch   bool
	verbose bool

	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	rootCmd := &cobra.Command{
		RunE:    run,
		Use:     "sync-dotenv",
		Short:   "sync-dotenv helps you to keep your .env.example in sync with your .env file",
		Long:    "sync-dotenv helps you to keep your .env.example in sync with your .env file",
		Version: fmt.Sprintf("%s (commit: %s, date: %s)", version, commit, date),
	}

	rootCmd.Flags().StringVar(&envFilePath, "env", ".env", "path to your env file")
	rootCmd.Flags().StringVar(&exampleFilePath, "example", ".env.example", "path to your example env file")
	rootCmd.Flags().StringVar(&basePath, "base", ".", "base path for all paths")

	rootCmd.Flags().BoolVarP(&watch, "watch", "w", false, "watch for file changes and update the example file automatically")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose messages")

	if err := rootCmd.Execute(); err != nil {
		if verbose {
			fmt.Printf("Error occured: %+v\n", err)
		} else {
			fmt.Printf("Error occured: %s\n", err)
		}

		os.Exit(1)
	}
}

func run(_ *cobra.Command, _ []string) error {
	envFilePath = filepath.Join(basePath, envFilePath)
	exampleFilePath = filepath.Join(basePath, exampleFilePath)

	if fileNotExist(envFilePath) {
		return errors.Errorf("env file (%s) does not exist", envFilePath)
	}

	if watch {
		return watchFile()
	} else {
		return processFile()
	}
}

func watchFile() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return errors.Wrap(err, "could not create new file watcher")
	}
	defer watcher.Close()

	errChan := make(chan error)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Name == filepath.Base(envFilePath) &&
					(event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Create == fsnotify.Create) {
					if err := processFile(); err != nil {
						errChan <- errors.Wrap(err, "failed to process file while watching for changes")
						return
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}

				errChan <- errors.Wrap(err, "failed to watch for changes")
				return
			}
		}
	}()

	err = watcher.Add(basePath)
	if err != nil {
		return errors.Wrap(err, "could not watch path")
	}
	return <-errChan
}

func processFile() error {
	envFile, err := os.Open(envFilePath)
	if err != nil {
		return errors.Wrap(err, "could not open env file")
	}
	defer envFile.Close()

	entries := map[string]string{}
	if !fileNotExist(exampleFilePath) {
		entries, err = getEntriesFromFile(exampleFilePath)
		if err != nil {
			return errors.Wrap(err, "could not get entries from example file")
		}
	}

	exampleEnvFile, err := os.Create(exampleFilePath)
	if err != nil {
		return errors.Wrap(err, "could not open example env file")
	}
	defer exampleEnvFile.Close()

	scanner := bufio.NewScanner(envFile)
	writer := bufio.NewWriter(exampleEnvFile)
	defer writer.Flush()

	if err := mirror(scanner, writer, entries); err != nil {
		return errors.Wrap(err, "failed while mirroring files")
	}

	return nil
}

func mirror(scanner *bufio.Scanner, writer *bufio.Writer, entries map[string]string) error {
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		var writeText string
		if len(text) > 0 && !strings.HasPrefix(text, "#") {
			splits := strings.SplitN(text, "=", 2)
			writeText = splits[0] + "="

			if val, ok := entries[splits[0]]; ok {
				writeText += val
			}
		} else {
			writeText = text
		}

		if _, err := fmt.Fprintln(writer, writeText); err != nil {
			return errors.Wrap(err, "failed while writing to example file")
		}
	}

	if err := scanner.Err(); err != nil {
		return errors.Wrap(err, "failed while reading env file")
	}

	return nil
}

func getEntriesFromFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "could not open example file")
	}
	defer file.Close()

	m := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		if len(text) <= 0 || strings.HasPrefix(text, "#") {
			continue
		}

		splits := strings.SplitN(text, "=", 2)
		key, value := splits[0], splits[1]

		m[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.Wrap(err, "failed while reading env file")
	}

	return m, nil
}

func fileNotExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

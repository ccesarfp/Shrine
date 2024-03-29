package main

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var (
	show     bool
	filePath string
)

// createKey - represents the createKey command
func createKey() *cobra.Command {
	createKeyCmd := &cobra.Command{
		Use:   "create:key",
		Short: "Creates an ID to be used in Shrine encryption",
		Long: `Creates an ID to be used for Shrine encryption using UUID v7.

The new key will be automatically included as an environment variable or in the .env file.`,
		Run: func(cmd *cobra.Command, args []string) {

			search := "HAS_ENV_VARS"
			hasEnvVars := os.Getenv(search)

			// Creating UUID v7
			v7, err := uuid.NewV7()
			if err != nil {
				log.Fatal(err)
			}
			if show {
				fmt.Println("Your UUID is:", v7)
			}

			// If dont use Environment Vars
			if hasEnvVars == "" {
				isInFile := false
				envVar := search + "=" + v7.String()

				// Verifying if file exists
				f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
				if err != nil {
					log.Fatalln(err)
				}
				defer func(f *os.File) {
					_ = f.Close()
				}(f)

				// Opening file
				input, err := os.ReadFile(filePath)
				if err != nil {
					log.Fatalln(err)
				}

				// Verifying file
				lines := strings.Split(string(input), "\n")
				for i, line := range lines {
					if strings.Contains(line, search) {
						lines[i] = envVar
						isInFile = true
					}
				}

				// If OT_SECRET_KEY doest not exists, create with uuid
				if !isInFile {
					linesNew := append(lines, make([]string, 2)...)
					linesNew[len(lines)] = envVar
					lines = linesNew
				}

				// Including new token
				output := strings.Join(lines, "\n")
				err = os.WriteFile(filePath, []byte(output), 0644)
				if err != nil {
					log.Fatalln(err)
				}
			}

			// If use Environment Vars
			if hasEnvVars != "" {
				err := os.Setenv(search, v7.String())
				if err != nil {
					log.Fatalln(err)
				}
			}

		},
	}

	createKeyCmd.Flags().BoolVarP(&show, "show", "s", false, "displays the created value for the UUID in the terminal")
	createKeyCmd.Flags().StringVarP(&filePath, "file", "f", "./.env", "environment variables file path")
	return createKeyCmd
}

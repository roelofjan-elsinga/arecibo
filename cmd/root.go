package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var srcFile string
var targetFile string
var writeMode string
var terminalOutput bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "arecibo",
	Short: "Arecibo: Simple text randomizer.",
	Long:  `Transform text templates into randomized texts. Useful for SEO purposes and many similar texts.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if srcFile == "" {
			return errors.New("source file is empty")
		}

		if !terminalOutput && targetFile == "" {
			return errors.New("target file path is empty")
		}

		workDir, err := os.Getwd()

		if err != nil {
			return err
		}

		if srcFile[0:1] != "/" {
			srcFile = fmt.Sprintf("%s/%s", workDir, srcFile)
		}

		if !terminalOutput && targetFile[0:1] != "/" {
			targetFile = fmt.Sprintf("%s/%s", workDir, targetFile)
		}

		if _, err := os.Stat(srcFile); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("no such file: %s", srcFile)
		}

		wordMapping := map[string]string{}

		var randomizedContent string

		if source, err := os.Open(srcFile); err != nil {
			return err
		} else {
			var buf bytes.Buffer

			// Copy file contents to string

			if _, err := io.Copy(&buf, source); err != nil {
				return err
			}

			randomizedContent = buf.String()

			// Find all strings starting with { and ending with }

			r := regexp.MustCompile(`{([^{}]*)}`)
			matches := r.FindAllStringSubmatch(randomizedContent, -1)

			for _, m := range matches {

				if len(m) == 2 {

					rand.Seed(time.Now().UnixNano())

					// Get all available options seperated by |

					options := strings.Split(m[1], "|")

					// Pick a random index

					optionIndex := rand.Intn(len(options) - 1)

					// Map the source to the target string

					wordMapping[m[0]] = strings.Trim(options[optionIndex], " ")
				}

			}

			// Replace the sources with the targets in the original text

			for k, v := range wordMapping {
				randomizedContent = strings.Replace(randomizedContent, k, v, 1)
			}

		}

		// Output

		if terminalOutput {
			fmt.Printf(randomizedContent)
			return nil
		}

		// Output to file

		if file, err := os.Create(targetFile); err != nil {
			return err
		} else {
			defer file.Close()

			if _, err := file.Write([]byte(randomizedContent)); err != nil {
				return err
			}

			if err := file.Sync(); err != nil {
				return err
			}
		}

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVarP(&srcFile, "source", "s", "", "Source file to generate text from")
	rootCmd.Flags().StringVarP(&targetFile, "output", "o", "", "Target file to write text to")
	rootCmd.Flags().BoolVarP(&terminalOutput, "terminal", "t", false, "Whether to output the text to the terminal")
}

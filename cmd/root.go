package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mvrcoag/service-monitor/storage"
	"github.com/spf13/cobra"
)

type Report struct {
	Url        string
	StatusCode int
	Duration   time.Duration
}

func Execute() {
	var rootCmd = &cobra.Command{}

	rootCmd.AddCommand(AddCmd)
	rootCmd.AddCommand(ListCmd)
	rootCmd.AddCommand(DropCmd)
	rootCmd.AddCommand(ReportCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a URL to the system monitor",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must provide an URL")
			os.Exit(1)
		}

		url := args[0]

		var s storage.Storage

		storage.ReadStorage(&s)

		urlExists := false

		for i := 0; i < len(s.Urls); i++ {
			if url == s.Urls[i] {
				urlExists = true
			}
		}

		if urlExists {
			print("Url already exists")
			os.Exit(1)
		}

		s.Urls = append(s.Urls, url)

		storage.WriteStorage(&s)

		fmt.Println("Service " + url + " added")
	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the services",
	Run: func(cmd *cobra.Command, args []string) {
		var s storage.Storage

		storage.ReadStorage(&s)

		fmt.Println("Your services:")
		fmt.Println("")

		for i := 0; i < len(s.Urls); i++ {
			index := strconv.Itoa(i)
			fmt.Println("[" + index + "] " + s.Urls[i])
		}
	},
}

var DropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop a service by index",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("You must provide an index to drop")
			os.Exit(1)
		}

		index, err := strconv.Atoi(args[0])

		if err != nil {
			fmt.Println("The index is not valid")
			os.Exit(1)
		}

		var s storage.Storage

		storage.ReadStorage(&s)

		if index > len(s.Urls) {
			fmt.Println("The given index is not in the available range")
			os.Exit(1)
		}

		service := s.Urls[index]

		s.Urls = append(s.Urls[:index], s.Urls[index+1:]...)

		storage.WriteStorage(&s)

		fmt.Println("Service " + service + " droped")
	},
}

var ReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Give a report of the registered URLs",
	Run: func(cmd *cobra.Command, args []string) {
		var s storage.Storage

		storage.ReadStorage(&s)

		for i := 0; i < len(s.Urls); i++ {
			url := s.Urls[i]
			report, err := generateReport(url)

			if err != nil {
				fmt.Println("Error on " + url)
				continue
			}

			statusCode := strconv.Itoa(report.StatusCode)
			time := report.Duration.String()

			fmt.Println(url)
			fmt.Println("Status Code: " + statusCode)
			fmt.Println("Duration: " + time)
			fmt.Println("")
		}
	},
}

func generateReport(url string) (Report, error) {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		return Report{}, err
	}

	defer res.Body.Close()

	return Report{
		Url:        url,
		StatusCode: res.StatusCode,
		Duration:   time.Since(start),
	}, nil
}

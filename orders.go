package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// ProductOrder describes a product order.
type ProductOrder struct {
	Date    string
	Amount  string
	Product string
}

func main() {
	path := "src/github.com/javierjmc/groceries-purchase-frequency/sources/"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	productAmounts := make(map[string]float64)
	productPurchaseFrequency := make(map[string]float64)

	for _, file := range files {
		fmt.Println("Processing >> " + file.Name() + "\n")

		lines, err := readCsv(path + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, line := range lines {
			order := ProductOrder{
				Date:    line[0],
				Amount:  line[1],
				Product: line[2],
			}

			amount, err := strconv.ParseFloat(order.Amount, 64)
			if err != nil {
				log.Fatal(err)
			}

			productAmounts[order.Product] += amount
			productPurchaseFrequency[order.Product]++
		}
	}

	firstDate := strings.Split(files[0].Name(), ".")
	lastDate := strings.Split(files[len(files)-1].Name(), ".")
	timeIntervalInDays := timeDiff(firstDate[0], lastDate[0])
	fmt.Println("timeIntervalInDays >> ", timeIntervalInDays, " days\n")

	writeCsv(productPurchaseFrequency, productAmounts, timeIntervalInDays)
}

// readCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func readCsv(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// writeCsv accepts a map and generates a csv file from it.
func writeCsv(productPurchaseFrequency map[string]float64, productAmounts map[string]float64, timeIntervalInDays int) {
	fmt.Println("Creating output...")

	output := [][]string{}

	output = append(output, []string{"Product", "Purchasing Frequency (days)", "Avg. Amount"})

	for product, frequency := range productPurchaseFrequency {
		//fmt.Println("Product:", product, ", Frequency:", frequency, ", PF:", math.Round(float64(timeIntervalInDays)/frequency), "days", ", Avg. Amount:", math.Round(productAmounts[product]/frequency))

		purchasingFrequency := fmt.Sprintf("%f", math.Round(float64(timeIntervalInDays)/frequency))
		averageAmount := fmt.Sprintf("%f", math.Round(productAmounts[product]/frequency))

		output = append(output, []string{product, purchasingFrequency, averageAmount})
	}

	file, err := os.Create("product_purchase_frequency.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range output {
		err := writer.Write(value)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Output created!")
}

// timeDiff takes a two dates as string and returns the time difference in days
func timeDiff(initDate, endDate string) int {
	a, erra := time.Parse("2006-01-02", initDate)
	b, errb := time.Parse("2006-01-02", endDate)

	if erra != nil || errb != nil {
		log.Fatal("Couldn't parse times")
	}

	duration := b.Sub(a)
	return int(duration.Hours() / 24)
}

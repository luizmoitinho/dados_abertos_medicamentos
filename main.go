package main

import (
	"log"

	"github.com/luizmoitinho/trabalho_final_aaed/files"
	"github.com/luizmoitinho/trabalho_final_aaed/ordering"
)

const (
	dataset          = "dataset/dados_abertos_medicamentos.csv"
	dataset_minified = "dataset/dados_abertos_medicamentos_minified.csv"
)

func main() {
	log.Println("Starting ordering analysis")

	log.Println("Reading data from dataset file")
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}

	log.Println("Ordering data from array")
	ordering.QuickSort(&medicines, 0, len(medicines)-1)

	log.Println("Saving ordered data into another CSV file")
	files.SaveCSV(&medicines, "results/quicksort.csv")

	log.Println("Stopping ordering analysis")
}

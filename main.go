package main

import (
	"log"

	"github.com/luizmoitinho/trabalho_final_aaed/evaluation"
	"github.com/luizmoitinho/trabalho_final_aaed/files"
	"github.com/luizmoitinho/trabalho_final_aaed/ordering"
)

const (
	dataset          = "dataset/dados_abertos_medicamentos.csv"
	dataset_minified = "dataset/dados_abertos_medicamentos_minified.csv"
)

func main() {
	evaluation.QuickSort()
}

func run() {
	log.Println("Starting ordering analysis")

	log.Println("Reading data from dataset file")
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}

	log.Printf("Ordering data from array with %d length\n", len(medicines))
	ordering.ShellSort(&medicines, len(medicines))

	log.Println("Saving ordered data into another CSV file")
	files.SaveCSV(&medicines, "results/shellsort.csv")

	log.Println("Stopping ordering analysis")
}

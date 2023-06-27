package evaluation

import (
	"fmt"
	"time"

	"github.com/luizmoitinho/trabalho_final_aaed/files"
	"github.com/luizmoitinho/trabalho_final_aaed/ordering"
)

func QuickSortGoodCase() {
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}
	ordering.QuickSort(medicines, 0, len(medicines)-1)

	start := time.Now()
	ordering.QuickSort(medicines, 0, len(medicines)-1)
	spent := time.Since(start)
	fmt.Printf("EvaluationQuickSortGoodCase[%d] %s\n", 0, spent)

	fmt.Printf("Comparacoes: %d\n", ordering.Comparacoes)
	fmt.Printf("troca: %d\n", ordering.Troca)
}

func QuickSortRandomCase() {
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	ordering.QuickSort(medicines, 0, len(medicines)-1)
	spent := time.Since(start)
	fmt.Printf("EvaluationQuickSortRandomCase[%d] %s\n", 0, spent)

	fmt.Printf("Comparacoes: %d\n", ordering.Comparacoes)
	fmt.Printf("troca: %d\n", ordering.Troca)
}

func QuickSortWorstCase() {
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}
	ordering.QuickSortDesc(medicines, 0, len(medicines)-1)

	start := time.Now()
	ordering.QuickSort(medicines, 0, len(medicines)-1)
	spent := time.Since(start)
	fmt.Printf("EvaluationQuickSortWorstCase[%d] %s\n", 0, spent)

	fmt.Printf("Comparacoes: %d\n", ordering.Comparacoes)
	fmt.Printf("troca: %d\n", ordering.Troca)
}

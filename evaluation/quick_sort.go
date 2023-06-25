package evaluation

import (
	"fmt"
	"time"

	"github.com/luizmoitinho/trabalho_final_aaed/entity"
	"github.com/luizmoitinho/trabalho_final_aaed/files"
	"github.com/luizmoitinho/trabalho_final_aaed/ordering"
)

func QuickSortGoodCase(n int) {
	fmt.Println("\nBegin of evaluation quicksort good case")
	data := make([]entity.Medicines, 0)
	for i := 0; i < 10; i++ {
		medicines, err := files.ReadCSV(dataset)
		if err != nil {
			panic(err)
		}
		ordering.QuickSort(&medicines, 0, len(medicines)-1)
		data = append(data, medicines)
	}

	for i := 0; i < n; i++ {
		start := time.Now()
		ordering.QuickSort(&data[i], 0, len(data[i])-1)
		fmt.Printf("EvaluationQuickSortGoodCase[%d] %s\n", i, time.Now().Sub(start))
	}
}

func QuickSortRandomCase(n int) {
	fmt.Println("\nBegin of evaluation quicksort random case")
	data := make([]entity.Medicines, 0)
	for i := 0; i < 10; i++ {
		medicines, err := files.ReadCSV(dataset)
		if err != nil {
			panic(err)
		}
		data = append(data, medicines)
	}

	for i := 0; i < n; i++ {
		start := time.Now()
		ordering.QuickSort(&data[i], 0, len(data[i])-1)
		fmt.Printf("EvaluationQuickSortRandomCase[%d] %s\n", i, time.Now().Sub(start))
	}
}

func QuickSortWorstCase(n int) {
	fmt.Println("\nBegin of evaluation quicksort worst case")
	data := make([]entity.Medicines, 0)
	for i := 0; i < 10; i++ {
		medicines, err := files.ReadCSV(dataset)
		if err != nil {
			panic(err)
		}
		ordering.QuickSortDesc(&medicines, 0, len(medicines)-1)
		data = append(data, medicines)
	}

	for i := 0; i < n; i++ {
		start := time.Now()
		ordering.QuickSort(&data[i], 0, len(data[i])-1)
		spent := time.Now().Sub(start)
		fmt.Printf("EvaluationQuickSortWorstCase[%d] %s\n", i, spent)
	}
}

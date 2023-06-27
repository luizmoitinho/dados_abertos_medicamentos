package evaluation

import (
	"fmt"
	"time"

	"github.com/luizmoitinho/trabalho_final_aaed/files"
	"github.com/luizmoitinho/trabalho_final_aaed/ordering"
)

func ShellSortGoodCase() {
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}
	ordering.ShellSort(medicines, len(medicines))

	start := time.Now()
	ordering.ShellSort(medicines, len(medicines))
	spent := time.Since(start)
	fmt.Printf("ShellSortGoodCase %s\n", spent)

	fmt.Printf("Comparacoes: %d\n", ordering.Comparacoes)
	fmt.Printf("troca: %d\n", ordering.Troca)
}

func ShellSortRandomCase() {
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	ordering.ShellSort(medicines, len(medicines))
	spent := time.Since(start)
	fmt.Printf("ShellSortRandomCase %s\n", spent)

	fmt.Printf("Comparacoes: %d\n", ordering.Comparacoes)
	fmt.Printf("troca: %d\n", ordering.Troca)
}

func ShellSortWorstCase() {
	medicines, err := files.ReadCSV(dataset)
	if err != nil {
		panic(err)
	}
	ordering.QuickSortDesc(medicines, 0, len(medicines)-1)

	start := time.Now()
	ordering.ShellSort(medicines, len(medicines))
	spent := time.Since(start)
	fmt.Printf("ShellSortWorstCase %s\n", spent)

	fmt.Printf("Comparacoes: %d\n", ordering.Comparacoes)
	fmt.Printf("troca: %d\n", ordering.Troca)
}

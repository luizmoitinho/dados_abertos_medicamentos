package evaluation

import (
	"fmt"
	"time"

	"github.com/luizmoitinho/trabalho_final_aaed/entity"
	"github.com/luizmoitinho/trabalho_final_aaed/files"
	"github.com/luizmoitinho/trabalho_final_aaed/ordering"
)

func ShellSortGoodCase(n int) {
	fmt.Println("\nBegin of evaluation shellSort good case")
	data := make([]entity.Medicines, 0)
	for i := 0; i < n; i++ {
		medicines, err := files.ReadCSV(dataset)
		if err != nil {
			panic(err)
		}
		ordering.ShellSort(&medicines, len(medicines))
		data = append(data, medicines)
	}

	for i := 0; i < n; i++ {
		start := time.Now()
		ordering.ShellSort(&data[i], len(data[i]))
		fmt.Printf("ShellSortGoodCase[%d] %s\n", i, time.Now().Sub(start))
	}
}

func ShellSortRandomCase(n int) {
	fmt.Println("\nBegin of evaluation shellSort random case")
	data := make([]entity.Medicines, 0)
	for i := 0; i < n; i++ {
		medicines, err := files.ReadCSV(dataset)
		if err != nil {
			panic(err)
		}
		data = append(data, medicines)
	}

	for i := 0; i < n; i++ {
		start := time.Now()
		ordering.ShellSort(&data[i], len(data[i]))
		fmt.Printf("ShellSortRandomCase[%d] %s\n", i, time.Now().Sub(start))
	}
}

func ShellSortWorstCase(n int) {
	fmt.Println("\nBegin of evaluation shellSort worst case")
	data := make([]entity.Medicines, 0)
	for i := 0; i < n; i++ {
		medicines, err := files.ReadCSV(dataset)
		if err != nil {
			panic(err)
		}
		ordering.QuickSortDesc(&medicines, 0, len(medicines)-1)
		data = append(data, medicines)
	}

	for i := 0; i < n; i++ {
		start := time.Now()
		ordering.ShellSort(&data[i], len(data[i]))
		spent := time.Now().Sub(start)
		fmt.Printf("ShellSortRandomCase[%d] %s\n", i, spent)
	}
}

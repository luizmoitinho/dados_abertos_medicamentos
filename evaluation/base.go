package evaluation

import "fmt"

var (
	dataset = "dataset/dados_abertos_medicamentos.csv"
)

func ShellSort() {
	ShellSortGoodCase()
	ShellSortGoodCase()
	ShellSortGoodCase()
	ShellSortGoodCase()
	ShellSortGoodCase()

	delimiter()

	ShellSortWorstCase()
	ShellSortWorstCase()
	ShellSortWorstCase()
	ShellSortWorstCase()
	ShellSortWorstCase()

	delimiter()

	ShellSortRandomCase()
	ShellSortRandomCase()
	ShellSortRandomCase()
	ShellSortRandomCase()
	ShellSortRandomCase()

	delimiter()
}

func QuickSort() {
	QuickSortGoodCase()
	QuickSortGoodCase()
	QuickSortGoodCase()
	QuickSortGoodCase()
	QuickSortGoodCase()

	delimiter()

	QuickSortWorstCase()
	QuickSortWorstCase()
	QuickSortWorstCase()
	QuickSortWorstCase()
	QuickSortWorstCase()

	delimiter()

	QuickSortRandomCase()
	QuickSortRandomCase()
	QuickSortRandomCase()
	QuickSortRandomCase()
	QuickSortRandomCase()

}

func delimiter() {
	fmt.Println("_______________________________")
}

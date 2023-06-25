package evaluation

var (
	dataset = "dataset/dados_abertos_medicamentos.csv"
)

func ShellSort() {
	ShellSortGoodCase(10)
	ShellSortRandomCase(10)
	ShellSortWorstCase(10)
}

func QuickSort() {
	QuickSortGoodCase(10)
	QuickSortRandomCase(10)
	QuickSortWorstCase(10)
}

package ordering

import (
	"github.com/luizmoitinho/trabalho_final_aaed/entity"
)

var (
	Comparacoes = int(0)
	Troca       = int(0)
)

func init() {
	Comparacoes = int(0)
	Troca = int(0)
}

func partition(medicines entity.Medicines, low, high int) int {
	pivot := medicines[high]
	i := low - 1
	for j := low; j < high; j++ {
		Comparacoes++
		if medicines[j].DueDate.Before(pivot.DueDate) {
			i++
			medicines[i], medicines[j] = medicines[j], medicines[i] //swap
			Troca++
		}
	}
	medicines[i+1], medicines[high] = medicines[high], medicines[i+1] //swap
	Troca++
	return i + 1
}

func QuickSort(medicines entity.Medicines, low, high int) {
	if low < high {
		pivot := partition(medicines, low, high)
		QuickSort(medicines, low, pivot-1)
		QuickSort(medicines, pivot+1, high)
	}
}

func partitionInverse(medicines entity.Medicines, low, high int) int {
	pivot := medicines[high]
	i := low
	for j := low; j < high; j++ {
		if medicines[j].DueDate.After(pivot.DueDate) {
			medicines[i], medicines[j] = medicines[j], medicines[i] //swap
			i++
		}
	}
	medicines[i], medicines[high] = medicines[high], medicines[i] //swap
	return i
}

func QuickSortDesc(medicines entity.Medicines, low, high int) {
	var pivot int
	if low < high {
		pivot = partitionInverse(medicines, low, high)
		QuickSortDesc(medicines, low, pivot-1)
		QuickSortDesc(medicines, pivot+1, high)
	}
}

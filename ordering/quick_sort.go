package ordering

import (
	"time"

	"github.com/luizmoitinho/trabalho_final_aaed/entity"
)

const dateFormat = "02/01/2006"

func partition(medicines *entity.Medicines, low, high int) int {
	pivot := (*medicines)[high]
	i := low
	for j := low; j < high; j++ {
		currentDueDate, _ := time.Parse(dateFormat, (*medicines)[j].DueDate)
		pivotDueDate, _ := time.Parse(dateFormat, pivot.DueDate)
		if currentDueDate.Before(pivotDueDate) {
			(*medicines)[i], (*medicines)[j] = (*medicines)[j], (*medicines)[i] //swap
			i++
		}
	}
	(*medicines)[i], (*medicines)[high] = (*medicines)[high], (*medicines)[i] //swap
	return i
}

func QuickSort(medicines *entity.Medicines, low, high int) {
	var pivot int
	if low < high {
		pivot = partition(medicines, low, high)
		QuickSort(medicines, low, pivot-1)
		QuickSort(medicines, pivot+1, high)
	}
}

func partitionInverse(medicines *entity.Medicines, low, high int) int {
	pivot := (*medicines)[high]
	i := low
	for j := low; j < high; j++ {
		currentDueDate, _ := time.Parse(dateFormat, (*medicines)[j].DueDate)
		pivotDueDate, _ := time.Parse(dateFormat, pivot.DueDate)

		if currentDueDate.After(pivotDueDate) {
			(*medicines)[i], (*medicines)[j] = (*medicines)[j], (*medicines)[i] //swap
			i++
		}
	}
	(*medicines)[i], (*medicines)[high] = (*medicines)[high], (*medicines)[i] //swap
	return i
}

func QuickSortDesc(medicines *entity.Medicines, low, high int) {
	var pivot int
	if low < high {
		pivot = partitionInverse(medicines, low, high)
		QuickSortDesc(medicines, low, pivot-1)
		QuickSortDesc(medicines, pivot+1, high)
	}
}

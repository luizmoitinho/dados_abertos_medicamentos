package ordering

import "github.com/luizmoitinho/trabalho_final_aaed/entity"

func partition(medicines *entity.Medicines, low, high int) (*entity.Medicines, int) {
	pivot := (*medicines)[high]
	i := low
	for j := low; j < high; j++ {
		if (*medicines)[j].Name < pivot.Name {
			(*medicines)[i], (*medicines)[j] = (*medicines)[j], (*medicines)[i]
			i++
		}
	}
	(*medicines)[i], (*medicines)[high] = (*medicines)[high], (*medicines)[i]
	return medicines, i
}

func QuickSort(medicines *entity.Medicines, low, high int) *entity.Medicines {
	if low < high {
		var p int
		medicines, p = partition(medicines, low, high)
		medicines = QuickSort(medicines, low, p-1)
		medicines = QuickSort(medicines, p+1, high)
	}
	return medicines
}

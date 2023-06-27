package ordering

import (
	"github.com/luizmoitinho/trabalho_final_aaed/entity"
)

func ShellSort(medicines entity.Medicines, n int) {
	var (
		gap = n / 2
	)

	for gap > 0 {
		for i := gap; i < n; i++ {
			aux := medicines[i]
			j := i
			Comparacoes++
			for j > gap-1 && (medicines[j-gap].DueDate.Before(aux.DueDate) || medicines[j-gap].DueDate.Equal(aux.DueDate)) {
				Troca++
				medicines[j] = medicines[j-gap]
				j -= gap
			}
			Troca++
			medicines[j] = aux
		}
		gap /= 2
	}
}

func ShellSortInverse(medicines entity.Medicines, n int) {
	var (
		gap = n / 2
	)

	for gap > 0 {
		for i := gap; i < n; i++ {
			aux := medicines[i]
			j := i
			for j >= gap && medicines[j-gap].DueDate.Before(aux.DueDate) {
				medicines[j] = medicines[j-gap]
				j -= gap
			}
			medicines[j] = aux

		}
		gap /= 2
	}

}

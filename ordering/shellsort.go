package ordering

import (
	"github.com/luizmoitinho/trabalho_final_aaed/entity"
)

func ShellSort(medicines *entity.Medicines, n int) {
	var (
		height int
	)

	for height < n/3 {
		height = 3*height + 1
	}

	for height > 0 {
		for i := height; i < n; i++ {
			j := i
			for j > height-1 && (*medicines)[j].GetTimeDueDate().Before((*medicines)[j-height].GetTimeDueDate()) {
				(*medicines)[j], (*medicines)[j-height] = (*medicines)[j-height], (*medicines)[j] //swap
				j -= height
			}
		}
		height /= 3
	}
}

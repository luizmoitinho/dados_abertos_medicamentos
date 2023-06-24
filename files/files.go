package files

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/luizmoitinho/trabalho_final_aaed/entity"
)

func ReadCSV(fileName string) (entity.Medicines, error) {
	var data entity.Medicines

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		item := entity.Medicine{
			Name:                       strings.TrimSpace(rec[0]),
			ProcessEndDate:             rec[1],
			RegulatoryCategory:         rec[2],
			DueDate:                    rec[4],
			ProcessNumber:              rec[5],
			TherapeuticClass:           rec[6],
			CompanyHoldingRegistration: rec[7],
			RegistrationStatus:         rec[8],
			ActiveIngredient:           rec[9],
		}

		value, err := strconv.ParseInt(rec[3], 10, 64)
		if err != nil {
			return nil, err
		}
		item.ProductRegistrtationNumber = value

		data = append(data, item)
	}
	return data, nil
}

func SaveCSV(medicines *entity.Medicines, fileName string) error {
	f, err := os.Create(fileName)

	if err != nil {
		return err
	}
	defer f.Close()

	for _, medicine := range *medicines {
		_, err = f.WriteString(fmt.Sprintf("%s,%s,%s,%d,%s,%s,%s,%s,%s,%s\n",
			medicine.Name,
			medicine.ProcessEndDate,
			medicine.RegulatoryCategory,
			medicine.ProductRegistrtationNumber,
			medicine.DueDate,
			medicine.ProcessNumber,
			medicine.TherapeuticClass,
			medicine.CompanyHoldingRegistration,
			medicine.RegistrationStatus,
			medicine.ActiveIngredient,
		))
		if err != nil {
			return err
		}
	}

	return nil
}

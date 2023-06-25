package entity

import "time"

/*
	NOME_PRODUTO
	DATA_FINALIZACAO_PROCESSO
	CATEGORIA_REGULATORIA
	NUMERO_REGISTRO_PRODUTO
	DATA_VENCIMENTO_REGISTRO
	NUMERO_PROCESSO
	CLASSE_TERAPEUTICA
	EMPRESA_DETENTORA_REGISTRO
	SITUACAO_REGISTRO
	PRINCIPIO_ATIVO
*/
const dateFormat = "02/01/2006"

type Medicines []Medicine

type Medicine struct {
	Name                       string
	ProcessEndDate             string
	RegulatoryCategory         string
	ProductRegistrtationNumber int64
	DueDate                    string
	ProcessNumber              string
	TherapeuticClass           string
	CompanyHoldingRegistration string
	RegistrationStatus         string
	ActiveIngredient           string
}

func (m Medicine) GetTimeDueDate() time.Time {
	dueDate, _ := time.Parse(dateFormat, m.DueDate)
	return dueDate
}

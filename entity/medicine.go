package entity

import "strings"

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

func (m *Medicine) NormalizeName() {
	m.Name = strings.TrimSpace(m.Name)
	m.Name = strings.Replace(m.Name, "\"", "", -1)
	m.Name = strings.Replace(m.Name, "(", "", -1)
	m.Name = strings.Replace(m.Name, ")", "", -1)

}

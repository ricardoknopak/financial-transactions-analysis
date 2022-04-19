package models

import (
	"gorm.io/gorm"
)

type Transactions struct {
	gorm.Model
	BancoOrigem    string  `json:"banco_origem"`
	AgenciaOrigem  string  `json:"agencia_origem"`
	ContaOrigem    string  `json:"conta_origem"`
	BancoDestino   string  `json:"banco_destino"`
	AgenciaDestino string  `json:"agencia_destino"`
	ContaDestino   string  `json:"conta_destino"`
	ValorTransacao float64 `json:"valor_transacao"`
}

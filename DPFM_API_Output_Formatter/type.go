package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey    string      `json:"connection_key"`
	Result           bool        `json:"result"`
	RedisKey         string      `json:"redis_key"`
	Filepath         string      `json:"filepath"`
	APIStatusCode    int         `json:"api_status_code"`
	RuntimeSessionID string      `json:"runtime_session_id"`
	BusinessPartner  int         `json:"business_partner"`
	ServiceLabel     string      `json:"service_label"`
	Message          Message     `json:"message"`
	APISchema        string      `json:"api_schema"`
	Accepter         []string    `json:"accepter"`
	OrderID          interface{} `json:"order_id"`
	Deleted          bool        `json:"deleted"`
}

type Message struct {
	Header Header `json:"Header"`
}

type Header struct {
	InvoiceDocument            int     `json:"InvoiceDocument"`
	CreationDate               string  `json:"CreationDate"`
	LastChangeDate             string  `json:"LastChangeDate"`
	BillToParty                int     `json:"BillToParty"`
	BillFromParty              int     `json:"BillFromParty"`
	BillToCountry              string  `json:"BillToCountry"`
	BillFromCountry            string  `json:"BillFromCountry"`
	InvoiceDocumentDate        string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime        string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate     string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate       string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate      string  `json:"AccountingPostingDate"`
	InvoiceDocumentIsCancelled bool    `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument   int     `json:"CancelledInvoiceDocument"`
	IsExportImportDelivery     bool    `json:"IsExportImportDelivery"`
	HeaderBillingIsConfirmed   bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus    string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount             float32 `json:"TotalNetAmount"`
	TotalTaxAmount             float32 `json:"TotalTaxAmount"`
	TotalGrossAmount           float32 `json:"TotalGrossAmount"`
	TransactionCurrency        string  `json:"TransactionCurrency"`
	Incoterms                  string  `json:"Incoterms"`
	PaymentTerms               string  `json:"PaymentTerms"`
	DueCalculationBaseDate     string  `json:"DueCalculationBaseDate"`
	NetPaymentDays             int     `json:"NetPaymentDays"`
	PaymentMethod              string  `json:"PaymentMethod"`
	HeaderPaymentBlockStatus   bool    `json:"HeaderPaymentBlockStatus"`
	ExternalReferenceDocument  string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText         string  `json:"DocumentHeaderText"`
}

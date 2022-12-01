package dpfm_api_input_reader

import (
	"convert-to-dpfm-invoice-document-from-sap-billing-document/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		BillingDocument:            data.BillingDocument,
		BillingDocumentType:        data.BillingDocumentType,
		SDDocumentCategory:         data.SDDocumentCategory,
		BillingDocumentCategory:    data.BillingDocumentCategory,
		CreationDate:               data.CreationDate,
		LastChangeDate:             data.LastChangeDate,
		SalesOrganization:          data.SalesOrganization,
		DistributionChannel:        data.DistributionChannel,
		Division:                   data.Division,
		BillingDocumentDate:        data.BillingDocumentDate,
		BillingDocumentIsCancelled: data.BillingDocumentIsCancelled,
		CancelledBillingDocument:   data.CancelledBillingDocument,
		IsExportDelivery:           data.IsExportDelivery,
		TotalNetAmount:             data.TotalNetAmount,
		TransactionCurrency:        data.TransactionCurrency,
		TaxAmount:                  data.TaxAmount,
		TotalGrossAmount:           data.TotalGrossAmount,
		CustomerPriceGroup:         data.CustomerPriceGroup,
		IncotermsClassification:    data.IncotermsClassification,
		CustomerPaymentTerms:       data.CustomerPaymentTerms,
		PaymentMethod:              data.PaymentMethod,
		CompanyCode:                data.CompanyCode,
		AccountingDocument:         data.AccountingDocument,
		ExchangeRateDate:           data.ExchangeRateDate,
		ExchangeRateType:           data.ExchangeRateType,
		DocumentReferenceID:        data.DocumentReferenceID,
		SoldToParty:                data.SoldToParty,
		PartnerCompany:             data.PartnerCompany,
		PurchaseOrderByCustomer:    data.PurchaseOrderByCustomer,
		CustomerGroup:              data.CustomerGroup,
		Country:                    data.Country,
		CityCode:                   data.CityCode,
		Region:                     data.Region,
		CreditControlArea:          data.CreditControlArea,
		OverallBillingStatus:       data.OverallBillingStatus,
		AccountingPostingStatus:    data.AccountingPostingStatus,
		AccountingTransferStatus:   data.AccountingTransferStatus,
		InvoiceListStatus:          data.InvoiceListStatus,
		BillingDocumentListType:    data.BillingDocumentListType,
		BillingDocumentListDate:    data.BillingDocumentListDate,
	}
}

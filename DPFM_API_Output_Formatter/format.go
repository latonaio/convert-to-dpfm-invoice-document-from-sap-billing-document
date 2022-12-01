package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-invoice-document-from-sap-billing-document/DPFM_API_Input_Reader"
	"fmt"
	"strconv"
)

func ConvertToHeader(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.Header
	var err error

	var totalNetAmountFloat32 *float32
	if data.TotalNetAmount != nil {
		totalNetAmountFloat32, err = parseFloat32Ptr(*data.TotalNetAmount)
		if err != nil {
			return nil, err
		}
	} else {
		totalNetAmountFloat32 = nil
	}

	var totalTaxAmountFloat32 *float32
	if data.TaxAmount != nil {
		totalTaxAmountFloat32, err = parseFloat32Ptr(*data.TaxAmount)
		if err != nil {
			return nil, err
		}
	} else {
		totalTaxAmountFloat32 = nil
	}

	header := Header{
		// InvoiceDocument:            data.InvoiceDocument,
		// CreationDate:               data.CreationDate,
		// LastChangeDate:             data.LastChangeDate,
		// BillToParty:                data.BillToParty,
		// BillFromParty: data.business_partner,
		// BillToCountry:              data.BillToCountry,
		BillFromCountry:     data.Country,
		InvoiceDocumentDate: data.BillingDocumentDate,
		// InvoiceDocumentTime:        data.InvoiceDocumentTime,
		// InvoicePeriodStartDate:     data.InvoicePeriodStartDate,
		// InvoicePeriodEndDate:       data.InvoicePeriodEndDate,
		// AccountingPostingDate:      data.AccountingPostingDate,
		InvoiceDocumentIsCancelled: data.BillingDocumentIsCancelled,
		// CancelledInvoiceDocument:   data.CancelledInvoiceDocument,
		// IsExportImportDelivery:     data.IsExportImportDelivery,
		// HeaderBillingIsConfirmed:   data.HeaderBillingIsConfirmed,
		// HeaderBillingConfStatus:    data.HeaderBillingConfStatus,
		TotalNetAmount: totalNetAmountFloat32,
		TotalTaxAmount: totalTaxAmountFloat32,
		// TotalGrossAmount:           data.TotalGrossAmount,
		// TransactionCurrency:        data.TransactionCurrency,
		// Incoterms:                  data.Incoterms,
		// PaymentTerms:               data.PaymentTerms,
		// DueCalculationBaseDate:     data.DueCalculationBaseDate,
		// NetPaymentDays:             data.NetPaymentDays,
		// PaymentMethod:              data.PaymentMethod,
		// HeaderPaymentBlockStatus:   data.HeaderPaymentBlockStatus,
		// ExternalReferenceDocument:  data.ExternalReferenceDocument,
		// DocumentHeaderText:         data.DocumentHeaderText,
	}

	return &header, nil
}

func parseFloat32Ptr(s string) (*float32, error) {
	resFloat64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil, err
	}
	resFloat32 := getFloat32Ptr(float32(resFloat64))

	return resFloat32, nil
}

func getFloat32Ptr(f float32) *float32 {
	return &f
}

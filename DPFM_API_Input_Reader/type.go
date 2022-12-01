package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey string   `json:"connection_key"`
	Result        bool     `json:"result"`
	RedisKey      string   `json:"redis_key"`
	Filepath      string   `json:"filepath"`
	Header        Header   `json:"Header"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	SONo          string   `json:"sales_order"`
	Deleted       bool     `json:"deleted"`
}

type Header struct {
	BillingDocument            *string `json:"BillingDocument"`
	BillingDocumentType        *string `json:"BillingDocumentType"`
	SDDocumentCategory         *string `json:"SDDocumentCategory"`
	BillingDocumentCategory    *string `json:"BillingDocumentCategory"`
	CreationDate               *string `json:"CreationDate"`
	LastChangeDate             *string `json:"LastChangeDate"`
	SalesOrganization          *string `json:"SalesOrganization"`
	DistributionChannel        *string `json:"DistributionChannel"`
	Division                   *string `json:"Division"`
	BillingDocumentDate        *string `json:"BillingDocumentDate"`
	BillingDocumentIsCancelled *bool   `json:"BillingDocumentIsCancelled"`
	CancelledBillingDocument   *string `json:"CancelledBillingDocument"`
	IsExportDelivery           *string `json:"IsExportDelivery"`
	TotalNetAmount             *string `json:"TotalNetAmount"`
	TransactionCurrency        *string `json:"TransactionCurrency"`
	TaxAmount                  *string `json:"TaxAmount"`
	TotalGrossAmount           *string `json:"TotalGrossAmount"`
	CustomerPriceGroup         *string `json:"CustomerPriceGroup"`
	IncotermsClassification    *string `json:"IncotermsClassification"`
	CustomerPaymentTerms       *string `json:"CustomerPaymentTerms"`
	PaymentMethod              *string `json:"PaymentMethod"`
	CompanyCode                *string `json:"CompanyCode"`
	AccountingDocument         *string `json:"AccountingDocument"`
	ExchangeRateDate           *string `json:"ExchangeRateDate"`
	ExchangeRateType           *string `json:"ExchangeRateType"`
	DocumentReferenceID        *string `json:"DocumentReferenceID"`
	SoldToParty                *string `json:"SoldToParty"`
	PartnerCompany             *string `json:"PartnerCompany"`
	PurchaseOrderByCustomer    *string `json:"PurchaseOrderByCustomer"`
	CustomerGroup              *string `json:"CustomerGroup"`
	Country                    *string `json:"Country"`
	CityCode                   *string `json:"CityCode"`
	Region                     *string `json:"Region"`
	CreditControlArea          *string `json:"CreditControlArea"`
	OverallBillingStatus       *string `json:"OverallBillingStatus"`
	AccountingPostingStatus    *string `json:"AccountingPostingStatus"`
	AccountingTransferStatus   *string `json:"AccountingTransferStatus"`
	InvoiceListStatus          *string `json:"InvoiceListStatus"`
	BillingDocumentListType    *string `json:"BillingDocumentListType"`
	BillingDocumentListDate    *string `json:"BillingDocumentListDate"`
}

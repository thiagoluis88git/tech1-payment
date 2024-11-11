package dto

type ExternalPaymentInformation struct {
	ID                int64
	Status            string
	ExternalReference string
	PreferenceID      string
	Marketplace       string
	NotificationURL   string
	DateCreated       string
	LastUpdated       string
	OrderStatus       string
	ClientID          string
}

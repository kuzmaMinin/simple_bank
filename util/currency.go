package util

const (
	USD = "USD"
	EUR = "EUR"
	RUR = "RUR"
)

func IsSupportedCurrency(curr string) bool {
	switch curr {
	case USD, EUR, RUR:
		return true
	}
	return false
}
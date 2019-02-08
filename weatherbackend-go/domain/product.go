package domain

type Product int

const (
	ForecastBasic = 0

	ForecastProfessional = 1

	ForecastUltimateXL = 2
)

func (product Product) String() string {
	names := [...]string{
		"ForecastBasic",
		"ForecastProfessional",
		"ForecastUltimateXL"}

	if product < ForecastBasic || product > ForecastUltimateXL {
		return "Unknown"
	}

	return names[product]
}

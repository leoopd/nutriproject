package main

type ScoreType int

const (
	Food ScoreType = iota
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	Value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

type EnergyKJ float64
type SugarGram float64
type SaturatedFattyAcids float64
type SodiumMilligram float64
type FruitsPercent float64
type FibreGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy              EnergyKJ
	Sugars              SugarGram
	SaturatedFattyAcids SaturatedFattyAcids
	Sodium              SodiumMilligram
	Fruits              FruitsPercent
	Fibre               FibreGram
	Protein             ProteinGram
	IsWater             bool
}

func (e EnergyKJ) GetPoints(st ScoreType) int {

}

func (s SugarGram) GetPoints(st ScoreType) int {

}

func (sfa SaturatedFattyAcids) GetPoints(st ScoreType) int {

}

func (so SodiumMilligram) GetPoints (st ScoreType) int {

}

func (f FruitsPercent)GetPoints (st ScoreType) int {

}

func (fg FibreGram) GetPoints (st ScoreType) int {

}

func  (p ProteinGram) GetPoints (st ScoreType) int {

}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {

	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints()
		fibrePoints := n.Fibre.GetPoints()

		negative = n.Energy.GetPoints() + n.Sugars.GetPoints() + n.SaturatedFattyAcids.GetPoints() + n.Sodium.GetPoints()
		positive = fruitPoints + fibrePoints + n.Protein.GetPoints()
	}

	return NutritionalScore{
		Value:     value,
		Positive:  positive,
		Negative:  negative,
		ScoreType: st,
	}
}

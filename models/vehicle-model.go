package model

type Vehicle struct {
	Id                                     string  `json:"id"`
	Make                                   string  `json:"Make"`
	Model                                  string  `json:"Model"`
	BaseModel                              string  `json:"baseModel"`
	Year                                   string  `json:"Year"`
	Drive                                  string  `json:"Drive"`
	FuelType                               string  `json:"Fuel Type"`
	FuelType1                              string  `json:"Fuel Type1"`
	FuelType2                              string  `json:"Fuel Type2"`
	Transmission                           string  `json:"Fuel Transmission"`
	TransmissionDescriptor                 string  `json:"Transmission descriptor"`
	EngineDescriptor                       string  `json:"Engine descriptor"`
	AnnualPetroleumConsumptionForFuelType1 float64 `json:"Annual Petroleum Consumption For Fuel Type1"`
	AnnualPetroleumConsumptionForFuelType2 float64 `json:"Annual Petroleum Consumption For Fuel Type2"`
	TimeToChargeAt120V                     float64 `json:"Time to charge at 120V"`
	TimeToChargeAt240V                     float64 `json:"Time to charge at 240V"`
	CityMpgForFuelType1                    float64 `json:"City Mpg For Fuel Type1"`
	CityMpgForFuelType2                    float64 `json:"City Mpg For Fuel Type2"`
	CityGasolineConsumption                float64 `json:"City gasoline consumption"`
	CityElectricityConsumption             float64 `json:"City electricity consumption"`
	EPACityUtilityFactor                   float64 `json:"EPA city utility factor"`
	EpaRangeForFuelType2                   float64 `json:"Epa Range For Fuel Type2"`
	EPAModelTypeIndex                      float64 `json:"EPA model type index"`
	EPAFuelEconomyScore                    float64 `json:"EPA Fuel Economy Score"`
	EPAHighwayUtilityFactor                float64 `json:"EPA highway utility factor"`
	EPACombinedFactor                      float64 `json:"EPA combined utility factor"`
	Co2FuelType1                           float64 `json:"Co2 Fuel Type1"`
	Co2FuelType2                           float64 `json:"Co2 Fuel Type2"`
	Co2TailpipeForFuelType1                float64 `json:"Co2  Tailpipe For Fuel Type1"`
	Co2TailpipeForFuelType2                float64 `json:"Co2  Tailpipe For Fuel Type2"`
	CombinedMpgForFuelType1                float64 `json:"Combined Mpg For Fuel Type1"`
	CombinedMpgForFuelType2                float64 `json:"Combined Mpg For Fuel Type2"`
	UnroundedCityMpgForFuelType1           float64 `json:"Unrounded City Mpg For Fuel Type1"`
	UnroundedCityMpgForFuelType2           float64 `json:"Unrounded City Mpg For Fuel Type2"`
	UnroundedCombinedMpgForFuelType1       float64 `json:"Unrounded Combined Mpg For Fuel Type1"`
	UnroundedCombinedMpgForFuelType2       float64 `json:"Unrounded Combined Mpg For Fuel Type2"`
	CombinedElectricityConsumption         float64 `json:"Combined electricity consumption"`
	CombinedGasolineConsumption            float64 `json:"Combined gasoline consumption"`
	Cylinders                              float64 `json:"Cylinders"`
	EngineDisplacement                     float64 `json:"Engine displacement"`
	AnnualFuelCostForFuelType1             float64 `json:"Annual Fuel Cost For Fuel Type1"`
	AnnualFuelCostForFuelType2             float64 `json:"Annual Fuel Cost For Fuel Type2"`
	GHGScore                               float64 `json:"GHG Score"`
	GHGScoreAlternativeFuel                float64 `json:"GHG Score Alternative Fuel"`
	HighwayMpgForFuelType1                 float64 `json:"Highway Mpg For Fuel Type1"`
	HighwayMpgForFuelType2                 float64 `json:"Highway Mpg For Fuel Type2"`
	UnroundedHighwayMpgForFuelType1        float64 `json:"Unrounded Highway Mpg For Fuel Type1"`
	UnroundedHighwayMpgForFuelType2        float64 `json:"Unrounded Highway Mpg For Fuel Type2"`
	HighwayGasolineConsumption             float64 `json:"Highway gasoline consumption"`
	HighwayElectricityConsumption          float64 `json:"Highway electricity consumption"`
	HatchbackLuggageVolume                 float64 `json:"Hatchback luggage volume"`
	HatchbackPassengerVolume               float64 `json:"Hatchback passenger volume"`
	Car2DoorLuggageVolume                  float64 `json:"Car 2 door luggage volume"`
	Car4DoorLuggageVolume                  float64 `json:"Car 4 door luggage volume"`
	Car2DoorPassengerVolume                float64 `json:"Car 2 door passenger volume"`
	Car4DoorPassengerVolume                float64 `json:"Car 4 door passenger volume"`
	MPGData                                string  `json:"MPG Data"`
	PHEVBlended                            bool    `json:"PHEV Blended"`
	RangeForFuelType1                      float64 `json:"Range For Fuel Type1"`
	RangeCityForFuelType1                  float64 `json:"Range  City For Fuel Type1"`
	RangeCityForFuelType2                  float64 `json:"Range  City For Fuel Type2"`
	RangeHighwayForFuelType1               float64 `json:"Range  Highway For Fuel Type1"`
	RangeHighwayForFuelType2               float64 `json:"Range  Highway For Fuel Type2"`
	UnadjustedCityMpgForFuelType1          float64 `json:"Unadjusted City Mpg For Fuel Type1"`
	UnadjustedCityMpgForFuelType2          float64 `json:"Unadjusted City Mpg For Fuel Type2"`
	UnadjustedHighwayMpgForFuelType1       float64 `json:"Unadjusted Highway Mpg For Fuel Type1"`
	UnadjustedHighwayMpgForFuelType2       float64 `json:"Unadjusted Highway Mpg For Fuel Type2"`
	Guzzler                                string  `json:"Guzzler"`
	TCharger                               string  `json:"T Charger"`
	SCharger                               string  `json:"S Charger"`
	ATVType                                string  `json:"ATV Type"`
	ElectricMotor                          string  `json:"Electric motor"`
	MFRCode                                string  `json:"MFR Code"`
	StartStop                              string  `json:"Start-Stop"`
	PHEVCity                               float64 `json:"PHEV City"`
	PHEVHighway                            float64 `json:"PHEV Highway"`
	PHEVCombined                           float64 `json:"PHEV Combined"`
}

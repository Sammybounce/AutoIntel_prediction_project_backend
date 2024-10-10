package controller

import (
	"ai-project/cache"
	Database "ai-project/utils/database"
	errorLogger "ai-project/utils/error-logger"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/google/uuid"
)

var ENV = os.Getenv("APP_ENV")

func PresetDefaults() {

	trainModel("decision-tree-model-training")

	trainModel("random-forest-model-training")

	// addVehicles()
}

func trainModel(notebook string) {

	if _, err := os.ReadFile(fmt.Sprintf("trained-data/%v.joblib", notebook)); err != nil {

		var trained bool
		var cmdStr string

		if runtime.GOOS == "windows" {
			cmdStr = fmt.Sprintf("jupyter nbconvert --to notebook --execute notebooks\\%v.ipynb --stdout", notebook)
		} else {
			cmdStr = fmt.Sprintf("jupyter nbconvert --to notebook --execute notebooks/%v.ipynb --stdout", notebook)
		}

		cmd := exec.Command(strings.Split(cmdStr, " ")[0], strings.Split(cmdStr, " ")[1:]...)

		stdout, err := cmd.Output()
		if err != nil {
			log.Fatalf("Failed to execute command: %s", err)
		}

		if strings.Contains(string(stdout), "success") {
			trained = true
		}

		if !trained {
			log.Fatal("Failed to train model")
		}
	}
}

func addVehicles() {
	db := Database.Connect()
	defer db.Close()

	vehicles := cache.VehicleCache()

	execStatement := `
		INSERT INTO vehicles
			(
				id,
				make,
				model,
				base_model,
				combined_name,
				year,
				drive,
				fuel_type,
				fuel_type_1,
				fuel_type_2,
				transmission,
				transmission_descriptor,
				engine_descriptor,
				annual_petroleum_consumption_for_fuel_type_1,
				annual_petroleum_consumption_for_fuel_type_2,
				time_to_charge_at_120V,
				time_to_charge_at_240V,
				city_mpg_for_fuel_type_1,
				city_mpg_for_fuel_type_2,
				city_gasoline_consumption,
				city_electricity_consumption,
				epa_city_utility_factor,
				epa_range_for_fuel_type_2,
				epa_model_type_index,
				epa_fuel_economy_score,
				epa_highway_utility_factor,
				epa_combined_utility_factor,
				co2_fuel_type_1,
				co2_fuel_type_2,
				co2_tailpipe_for_fuel_type_1,
				co2_tailpipe_for_fuel_type_2,
				combined_mpg_for_fuel_type_1,
				combined_mpg_for_fuel_type_2,
				unrounded_city_mpg_For_fuel_type1,
				unrounded_city_mpg_For_fuel_type2,
				unrounded_combined_mpg_For_fuel_type1,
				unrounded_combined_mpg_For_fuel_type2,
				combined_electricity_consumption,
				combined_gasoline_consumption,
				cylinders,
				engine_displacement,
				annual_fuel_cost_for_fuel_type_1,
				annual_fuel_cost_for_fuel_type_2,
				ghg_score,
				ghg_score_alternative_fuel,
				highway_mpg_for_fuel_type_1,
				highway_mpg_for_fuel_type_2,
				unrounded_highway_mpg_for_fuel_type_1,
				unrounded_highway_mpg_for_fuel_type_2,
				highway_electricity_consumption,
				highway_gasoline_consumption,
				hatchback_luggage_volume,
				hatchback_passenger_volume,
				car_2_door_luggage_volume,
				car_4_door_luggage_volume,
				car_2_door_passenger_volume,
				car_4_door_passenger_volume,
				mpg_data,
				phev_blended,
				range_for_fuel_type_1,
				range_city_for_fuel_type_1,
				range_city_for_fuel_type_2,
				range_highway_for_fuel_type_1,
				range_highway_for_fuel_type_2,
				unadjusted_city_mpg_for_fuel_type_1,
				unadjusted_city_mpg_for_fuel_type_2,
				unadjusted_highway_mpg_for_fuel_type_1,
				unadjusted_highway_mpg_for_fuel_type_2,
				guzzler,
				t_charger,
				s_charger,
				atv_ype,
				electric_motor,
				mfr_code,
				start_stop,
				phev_city,
				phev_highway,
				phev_combined
			)

		VALUES
			--VALUES ;
	`

	var values string

	for _, v := range *vehicles {

		rows, err := db.Query(fmt.Sprintf(`
			SELECT 1 
			
			FROM 
				vehicles 
				
			WHERE 
				LOWER(combined_name) = LOWER('%v');
		`,
			fmt.Sprintf("%v %v %v", v.Make, v.Model, v.BaseModel),
		))

		if err != nil {
			errorLogger.CaptureException("addVehicles--1", fmt.Errorf("%v | error: %v", v, err))
			return
		}
		defer rows.Close()

		if rows.Next() {
			rows.Close()
			continue
		}

		rows.Close()

		Id, err := uuid.NewV7()
		if err != nil {
			return
		}

		value := fmt.Sprintf(`
			('%v', '%v', '%v', '%v', '%v', %v, '%v', '%v', '%v', '%v', '%v', '%v', '%v', %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, '%v', '%v', %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, '%v', %v, %v, %v, %v, %v, %v, %v, %v, %v, %v, '%v', '%v', '%v', '%v', '%v', '%v', '%v', %v, %v, %v)
		`,

			Id,
			v.Make,
			v.Model,
			v.BaseModel,
			fmt.Sprintf("%v %v %v", v.Make, v.Model, v.BaseModel),
			v.Year,
			v.Drive,
			v.FuelType,
			v.FuelType1,
			v.FuelType2,
			v.Transmission,
			v.TransmissionDescriptor,
			v.EngineDescriptor,
			v.AnnualPetroleumConsumptionForFuelType1,
			v.AnnualPetroleumConsumptionForFuelType2,
			v.TimeToChargeAt120V,
			v.TimeToChargeAt240V,
			v.CityMpgForFuelType1,
			v.CityMpgForFuelType2,
			v.CityGasolineConsumption,
			v.CityElectricityConsumption,
			v.EPACityUtilityFactor,
			v.EpaRangeForFuelType2,
			v.EPAModelTypeIndex,
			v.EPAFuelEconomyScore,
			v.EPAHighwayUtilityFactor,
			v.EPACombinedFactor,
			v.Co2FuelType1,
			v.Co2FuelType2,
			v.Co2TailpipeForFuelType1,
			v.Co2TailpipeForFuelType2,
			v.CombinedElectricityConsumption,
			v.CombinedGasolineConsumption,
			v.UnroundedCityMpgForFuelType1,
			v.UnroundedCityMpgForFuelType2,
			v.UnroundedCombinedMpgForFuelType1,
			v.UnroundedCombinedMpgForFuelType2,
			v.CombinedElectricityConsumption,
			v.CombinedGasolineConsumption,
			v.Cylinders,
			v.EngineDisplacement,
			v.AnnualFuelCostForFuelType1,
			v.AnnualFuelCostForFuelType2,
			v.GHGScore,
			v.GHGScoreAlternativeFuel,
			v.HighwayMpgForFuelType1,
			v.HighwayMpgForFuelType2,
			v.UnroundedHighwayMpgForFuelType1,
			v.UnroundedHighwayMpgForFuelType2,
			v.HighwayElectricityConsumption,
			v.HighwayGasolineConsumption,
			v.HatchbackLuggageVolume,
			v.HatchbackPassengerVolume,
			v.Car2DoorLuggageVolume,
			v.Car4DoorLuggageVolume,
			v.Car2DoorPassengerVolume,
			v.Car4DoorPassengerVolume,
			v.MPGData,
			v.PHEVBlended,
			v.RangeForFuelType1,
			v.RangeCityForFuelType1,
			v.RangeCityForFuelType2,
			v.RangeHighwayForFuelType1,
			v.RangeHighwayForFuelType2,
			v.UnadjustedCityMpgForFuelType1,
			v.UnadjustedCityMpgForFuelType2,
			v.UnadjustedHighwayMpgForFuelType1,
			v.UnadjustedHighwayMpgForFuelType2,
			v.Guzzler,
			v.TCharger,
			v.SCharger,
			v.ATVType,
			v.ElectricMotor,
			v.MFRCode,
			v.StartStop,
			v.PHEVCity,
			v.PHEVHighway,
			v.PHEVCombined,
		)

		if values == "" {
			values = value
		} else {
			values = fmt.Sprintf(`
				%v, 
				%v
			`,
				values,
				value,
			)
		}
	}

	if values != "" {
		execStatement = strings.Replace(execStatement, "--VALUES", values, 1)

		tx, err := db.Begin()
		if err != nil {
			errorLogger.CaptureException("addVehicles--2", err)

			return
		}

		if _, err := tx.Exec(execStatement); err != nil {
			errorLogger.CaptureException("addVehicles--3", fmt.Errorf("%v | error: %v", execStatement, err))

			if err := tx.Rollback(); err != nil {
				errorLogger.CaptureException("addVehicles--4", err)

				return
			}

			return
		}

		if err := tx.Commit(); err != nil {
			errorLogger.CaptureException("addVehicles--5", err)

			return
		}
	}
}

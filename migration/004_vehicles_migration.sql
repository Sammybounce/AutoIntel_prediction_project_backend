-- +goose Up
-- all_car start
CREATE TABLE
  IF NOT EXISTS vehicles (
    id UUID NOT NULL,
    make VARCHAR(225) DEFAULT 'N/A',
    model VARCHAR(225) DEFAULT 'N/A',
    base_model VARCHAR(225) DEFAULT 'N/A',
    combined_name VARCHAR(225) DEFAULT 'N/A',
    year DOUBLE PRECISION DEFAULT 0,
    drive VARCHAR(225) DEFAULT 'N/A',
    fuel_type VARCHAR(225) DEFAULT 'N/A',
    fuel_type_1 VARCHAR(225) DEFAULT 'N/A',
    fuel_type_2 VARCHAR(225) DEFAULT 'N/A',
    transmission VARCHAR(225) DEFAULT 'N/A',
    transmission_descriptor VARCHAR(225) DEFAULT 'N/A',
    engine_descriptor VARCHAR(225) DEFAULT 'N/A',
    annual_petroleum_consumption_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    annual_petroleum_consumption_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    time_to_charge_at_120V DOUBLE PRECISION DEFAULT 0,
    time_to_charge_at_240V DOUBLE PRECISION DEFAULT 0,
    city_mpg_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    city_mpg_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    city_gasoline_consumption DOUBLE PRECISION DEFAULT 0,
    city_electricity_consumption DOUBLE PRECISION DEFAULT 0,
    epa_city_utility_factor DOUBLE PRECISION DEFAULT 0,
    epa_range_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    epa_model_type_index DOUBLE PRECISION DEFAULT 0,
    epa_fuel_economy_score DOUBLE PRECISION DEFAULT 0,
    epa_highway_utility_factor VARCHAR(225) DEFAULT 'N/A',
    epa_combined_utility_factor VARCHAR(225) DEFAULT 'N/A',
    co2_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    co2_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    co2_tailpipe_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    co2_tailpipe_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    combined_mpg_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    combined_mpg_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    unrounded_city_mpg_For_fuel_type1 DOUBLE PRECISION DEFAULT 0,
    unrounded_city_mpg_For_fuel_type2 DOUBLE PRECISION DEFAULT 0,
    unrounded_combined_mpg_For_fuel_type1 DOUBLE PRECISION DEFAULT 0,
    unrounded_combined_mpg_For_fuel_type2 DOUBLE PRECISION DEFAULT 0,
    combined_electricity_consumption DOUBLE PRECISION DEFAULT 0,
    combined_gasoline_consumption DOUBLE PRECISION DEFAULT 0,
    cylinders DOUBLE PRECISION DEFAULT 0,
    engine_displacement DOUBLE PRECISION DEFAULT 0,
    annual_fuel_cost_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    annual_fuel_cost_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    ghg_score DOUBLE PRECISION DEFAULT 0,
    ghg_score_alternative_fuel DOUBLE PRECISION DEFAULT 0,
    highway_mpg_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    highway_mpg_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    unrounded_highway_mpg_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    unrounded_highway_mpg_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    highway_electricity_consumption DOUBLE PRECISION DEFAULT 0,
    highway_gasoline_consumption DOUBLE PRECISION DEFAULT 0,
    hatchback_luggage_volume DOUBLE PRECISION DEFAULT 0,
    hatchback_passenger_volume DOUBLE PRECISION DEFAULT 0,
    car_2_door_luggage_volume DOUBLE PRECISION DEFAULT 0,
    car_4_door_luggage_volume DOUBLE PRECISION DEFAULT 0,
    car_2_door_passenger_volume DOUBLE PRECISION DEFAULT 0,
    car_4_door_passenger_volume DOUBLE PRECISION DEFAULT 0,
    mpg_data VARCHAR(4) DEFAULT 'N/A',
    phev_blended BOOLEAN DEFAULT false,
    range_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    range_city_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    range_city_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    range_highway_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    range_highway_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    unadjusted_city_mpg_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    unadjusted_city_mpg_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    unadjusted_highway_mpg_for_fuel_type_1 DOUBLE PRECISION DEFAULT 0,
    unadjusted_highway_mpg_for_fuel_type_2 DOUBLE PRECISION DEFAULT 0,
    guzzler VARCHAR(4) DEFAULT 'N/A',
    t_charger VARCHAR(4) DEFAULT 'N/A',
    s_charger VARCHAR(4) DEFAULT 'N/A',
    atv_ype VARCHAR(60) DEFAULT 'N/A',
    electric_motor VARCHAR(60) DEFAULT 'N/A',
    mfr_code VARCHAR(60) DEFAULT 'N/A',
    start_stop VARCHAR(4) DEFAULT 'N/A',
    phev_city DOUBLE PRECISION DEFAULT 0,
    phev_highway DOUBLE PRECISION DEFAULT 0,
    phev_combined DOUBLE PRECISION DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted BOOLEAN NOT NULL DEFAULT false,
    PRIMARY KEY (id)
  );

CREATE
OR REPLACE TRIGGER trigger_update_last_modified_on_vehicles BEFORE
UPDATE ON vehicles FOR EACH ROW EXECUTE PROCEDURE update_last_modified ();

-- all_car end
-- +goose Down
DROP TABLE IF EXISTS vehicles CASCADE;
package models

import (
	"errors"
	"time"

	"Conditions/db"
	forms "Conditions/forms"
)

//ConditionModel ...
type ConditionModel struct{}

var conditionModel = new(ConditionModel)

//GetTemperature ...
func (c ConditionModel) GetTemperature(deviceID string, timeframe int) (device forms.ReturnTemp, err error) {
	getDb := db.GetDB()
	checkDevice, err := getDb.SelectInt("Select count(id) from public.devices WHERE deviceid=LOWER($1) LIMIT 1",
		deviceID)
	if err != nil {
		return device, err
	}
	if checkDevice > 1 {
		return device, errors.New("More than one device was requested, please only request one Device")
	}
	if checkDevice < 1 {
		return device, errors.New("device does not exist")
	}
	if timeframe < 1 {
		_, err = getDb.Select(
			&device.Temperature,
			"SELECT temperature FROM public.conditions WHERE device=$1 ORDER BY \"time\" DESC",
			deviceID)
	} else {
		_, err = getDb.Select(
			&device.Temperature,
			"SELECT temperature FROM public.conditions WHERE device=$1 ORDER BY \"time\" DESC LIMIT $2",
			deviceID,
			timeframe)
	}
	i := make([]float64, 0, 10)
	for _, p := range device.Temperature {

		i = append(i, p)

	}
	device.Temperature = i
	return device, err
}

//SendTemperature Send the temperature to the Conditions table, expect no confirmation
func (c ConditionModel) SendTemperature(deviceID string, form forms.Conditions) (device forms.DeviceID, err error) {
	getDb := db.GetDB()
	temperature := form.Temperature
	humidity := form.Humidity
	err = getDb.QueryRow("INSERT INTO public.conditions(time, device, temperature, humidity) VALUES($1, $2, $3, $4) RETURNING device",
		time.Now(),
		deviceID,
		temperature,
		humidity).Scan(&device.DeviceID)

	return device, err

}

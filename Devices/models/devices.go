package models

import (
	"errors"

	db "Devices/db"
	"Devices/forms"
)

// DeviceModel ...
type DeviceModel struct{}

var devicemodel = new(DeviceModel)

//User ...
type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"-"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

// CreateNewDevice ...
func (m DeviceModel) CreateNewDevice(userID int64, form forms.Device) (device forms.Device, err error) {
	getDb := db.GetDB()
	checkDevice, err := getDb.SelectInt("Select count(id) from public.devices WHERE deviceid=LOWER($1) LIMIT 1", form.DeviceID)
	if err != nil {
		return device, err
	}

	if checkDevice > 0 {
		return device, errors.New("Device already exists")
	}
	Owner := User{}
	err = getDb.SelectOne(&Owner, "SELECT email, name FROM public.user WHERE id=$1", userID)
	if err != nil {
		return device, err
	}
	err = getDb.QueryRow("INSERT INTO public.devices(name, deviceid, owner) VALUES($1, $2, $3) RETURNING id",
		form.Name,
		form.DeviceID,
		Owner.ID).Scan(&device.ID)
	device.DeviceID = form.DeviceID
	device.Name = form.Name
	device.Owner = Owner.Email
	return device, err
}

// GetDeviceByID ...
func (m DeviceModel) GetDeviceByID(deviceID string) (device forms.Device, err error) {
	err = db.GetDB().SelectOne(&device, "Select id, name, deviceid, owner from public.devices where deviceid=$1", deviceID)

	return device, err
}

// GetDevicesByOwner ...
func (m DeviceModel) GetDevicesByOwner(userID int64) (device forms.DeviceID, err error) {
	getDb := db.GetDB()

	var posts []forms.Device
	_, err = getDb.Select(&posts, "Select * from public.devices where owner='1'")
	d := make([]string, 0, 10)
	for _, i := range posts {
		d = append(d, i.DeviceID)
	}
	device.DeviceID = d
	return device, err
}

// UpdateDevice ...
// TODO : Fix this so that it actually updates properly
func (m DeviceModel) UpdateDevice(form forms.Device) (device forms.Device, err error) {
	getDb := db.GetDB()

	count, err := getDb.SelectInt("SELECT id, name, deviceid, owner from public.devices where deviceid=$1", form.DeviceID)

	form.ID = count
	count, err = getDb.Update(device)
	return form, err
}

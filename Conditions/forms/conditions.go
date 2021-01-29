package forms

import "time"

// DeviceID ...
type DeviceID struct {
	DeviceID []string `json:"deviceid"`
}

//ReturnTemp ...
type ReturnTemp struct {
	Temperature []float64 `json:"temperature"`
}

// Conditions ...
type Conditions struct {
	time        time.Time `db:"time"`
	Device      string    `db:"device"`
	Humidity    float64   `db:"humidity" json:"humidity,omitempty"`
	Temperature float64   `db:"temperature" json:"temperature"`
}

//Temperature ...
type Temperature struct {
	Temperature float64 `db:"temperature" json:"temperature"`
}

//Humidity ...
type Humidity struct {
	Humidity float64 `db:"humidity" json:"humidity"`
}

// Device ...
type Device struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Name      string `db:"name" json:"name"`
	DeviceID  string `db:"deviceid" json:"deviceid"`
	Owner     string `db:"owner" json:"owner"`
	CreatedAt int64  `db:"created_at" json:"-"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
}

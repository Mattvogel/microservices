package forms

//DeviceID ...
type DeviceID struct {
	DeviceID []string `json:"deviceid"`
}

//UpdateDevice ...
type UpdateDevice struct {
	Name  string `json:"name"`
	Owner string `json:"owner,omitempty"`
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

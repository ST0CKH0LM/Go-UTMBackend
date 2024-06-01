package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type JSONTime time.Time

func (jt *JSONTime) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		*jt = JSONTime(v)
	case []byte:
		t, err := time.Parse(time.RFC3339, string(v))
		if err != nil {
			t, err = time.Parse("2006-01-02 15:04:05", string(v))
			if err != nil {
				return err
			}
		}
		*jt = JSONTime(t)
	case string:
		t, err := time.Parse(time.RFC3339, v)
		if err != nil {
			t, err = time.Parse("2006-01-02 15:04:05", string(v))
			if err != nil {
				return err
			}
		}
		*jt = JSONTime(t)
	default:
		return errors.New("unsupported data type for JSONTime")
	}
	return nil
}

func (jt JSONTime) Value() (driver.Value, error) {
	t := time.Time(jt)
	return t.Format("2006-01-02 15:04:05"), nil
}

func (jt JSONTime) MarshalJSON() ([]byte, error) {
	t := time.Time(jt)
	return json.Marshal(t)
}

func (jt *JSONTime) UnmarshalJSON(data []byte) error {
	var t time.Time
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	*jt = JSONTime(t)
	return nil
}

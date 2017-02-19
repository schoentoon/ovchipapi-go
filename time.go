package ovchipapi

import (
	"encoding/json"
	"time"
)

type OVTime time.Time

func (t OVTime) MarshalJSON() ([]byte, error) {
	return []byte(string(time.Time(t).Unix() * 1000)), nil
}

func (t *OVTime) UnmarshalJSON(b []byte) error {
	var millis int64
	if err := json.Unmarshal(b, &millis); err != nil {
		return err
	}
	*t = OVTime(time.Unix(0, millis*int64(time.Millisecond)))

	return nil
}

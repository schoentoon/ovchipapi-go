package ovchipapi

import (
	"encoding/json"
	"math"
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

	// we can throw away the millis part of the time
	// because the server just returns whatever
	// the current millis are, so we only have
	// accuracy to the second
	unix := int64(math.Floor(float64(millis) / 1000))

	*t = OVTime(time.Unix(unix, 0))

	return nil
}

func (t OVTime) String() string {
	return time.Time(t).String()
}

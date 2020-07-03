package utils

import (
	"encoding/json"
)

type Mask string

func (c *Mask) Equals(mask Mask) bool {
	return *c == mask
}

func (c Mask) MarshalJSON() ([]byte, error) {
	var asteriskString string
	for i := 0; i < len(c); i++ {
		asteriskString += "*"
	}
	return json.Marshal(asteriskString)
}

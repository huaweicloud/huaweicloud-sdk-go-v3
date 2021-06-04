package model

import (
	"encoding/json"

	"strings"
)

type EnlargeVolume struct {
	EnlargeVolume *EnlargeVolumeObject `json:"enlarge_volume"`
}

func (o EnlargeVolume) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnlargeVolume struct{}"
	}

	return strings.Join([]string{"EnlargeVolume", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

//
type OpenGaussInstanceActionRequest struct {
	ExpandCluster *OpenGaussExpandCluster `json:"expand_cluster,omitempty"`

	EnlargeVolume *OpenGaussEnlargeVolume `json:"enlarge_volume,omitempty"`
}

func (o OpenGaussInstanceActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OpenGaussInstanceActionRequest struct{}"
	}

	return strings.Join([]string{"OpenGaussInstanceActionRequest", string(data)}, " ")
}

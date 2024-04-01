package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NovaShowVersionResponse Response Object
type NovaShowVersionResponse struct {
	Version        *NovaVersionDetail `json:"version,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o NovaShowVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NovaShowVersionResponse struct{}"
	}

	return strings.Join([]string{"NovaShowVersionResponse", string(data)}, " ")
}

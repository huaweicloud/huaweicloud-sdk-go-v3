package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowVolumeResponse struct {
	Volume         *Volume `json:"volume,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVolumeResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeResponse", string(data)}, " ")
}

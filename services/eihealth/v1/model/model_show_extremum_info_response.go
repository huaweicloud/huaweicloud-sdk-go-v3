package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtremumInfoResponse Response Object
type ShowExtremumInfoResponse struct {
	Maximum *ExtremumDto `json:"maximum,omitempty"`

	Minimum        *ExtremumDto `json:"minimum,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowExtremumInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtremumInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowExtremumInfoResponse", string(data)}, " ")
}

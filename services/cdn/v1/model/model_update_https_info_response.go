package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateHttpsInfoResponse struct {
	Https          *HttpInfoResponseBody `json:"https,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateHttpsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpsInfoResponse struct{}"
	}

	return strings.Join([]string{"UpdateHttpsInfoResponse", string(data)}, " ")
}

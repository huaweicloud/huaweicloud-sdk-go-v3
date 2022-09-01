package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowHttpInfoResponse struct {
	Https          *HttpInfoResponseBody `json:"https,omitempty" xml:"https"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowHttpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpInfoResponse", string(data)}, " ")
}

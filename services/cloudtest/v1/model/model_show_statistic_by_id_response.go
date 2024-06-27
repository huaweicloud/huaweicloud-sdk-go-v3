package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatisticByIdResponse Response Object
type ShowStatisticByIdResponse struct {
	Code *string `json:"code,omitempty"`

	Data map[string]int32 `json:"data,omitempty"`

	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowStatisticByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatisticByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowStatisticByIdResponse", string(data)}, " ")
}

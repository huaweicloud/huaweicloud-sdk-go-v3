package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSendCountryDetailsResponse Response Object
type ListSendCountryDetailsResponse struct {
	Body           *[]SmsCountryResp `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListSendCountryDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSendCountryDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListSendCountryDetailsResponse", string(data)}, " ")
}

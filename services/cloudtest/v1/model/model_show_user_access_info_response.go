package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUserAccessInfoResponse Response Object
type ShowUserAccessInfoResponse struct {

	// 对外时：success|error;
	Status *string `json:"status,omitempty"`

	Result         *ResultValueStringForOk `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowUserAccessInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserAccessInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowUserAccessInfoResponse", string(data)}, " ")
}

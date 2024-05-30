package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmApprovalsResponse Response Object
type ConfirmApprovalsResponse struct {
	Data           *ConfirmApprovalsResultData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ConfirmApprovalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmApprovalsResponse struct{}"
	}

	return strings.Join([]string{"ConfirmApprovalsResponse", string(data)}, " ")
}

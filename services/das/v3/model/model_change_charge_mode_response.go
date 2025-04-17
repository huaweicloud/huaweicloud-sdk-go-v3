package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeChargeModeResponse Response Object
type ChangeChargeModeResponse struct {

	// 是否成功
	Success        *bool `json:"success,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ChangeChargeModeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeChargeModeResponse struct{}"
	}

	return strings.Join([]string{"ChangeChargeModeResponse", string(data)}, " ")
}

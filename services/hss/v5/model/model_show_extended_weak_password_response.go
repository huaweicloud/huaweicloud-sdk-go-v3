package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowExtendedWeakPasswordResponse Response Object
type ShowExtendedWeakPasswordResponse struct {

	// **参数解释**: 自定义弱口令，选填，可编辑 **取值范围**: 最小值0，最大值1000000
	ExtendedWeakPassword *[]string `json:"extended_weak_password,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o ShowExtendedWeakPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowExtendedWeakPasswordResponse struct{}"
	}

	return strings.Join([]string{"ShowExtendedWeakPasswordResponse", string(data)}, " ")
}

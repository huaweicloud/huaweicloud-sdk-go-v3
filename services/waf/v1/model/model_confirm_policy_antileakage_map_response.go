package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmPolicyAntileakageMapResponse Response Object
type ConfirmPolicyAntileakageMapResponse struct {
	Leakagemap *AntileakageMapResponseBodyLeakagemap `json:"leakagemap,omitempty"`

	Locale         *AntileakageMapResponseBodyLocale `json:"locale,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ConfirmPolicyAntileakageMapResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmPolicyAntileakageMapResponse struct{}"
	}

	return strings.Join([]string{"ConfirmPolicyAntileakageMapResponse", string(data)}, " ")
}

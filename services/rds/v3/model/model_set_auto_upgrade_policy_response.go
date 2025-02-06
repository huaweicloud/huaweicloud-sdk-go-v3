package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoUpgradePolicyResponse Response Object
type SetAutoUpgradePolicyResponse struct {

	// 响应结果
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAutoUpgradePolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoUpgradePolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAutoUpgradePolicyResponse", string(data)}, " ")
}

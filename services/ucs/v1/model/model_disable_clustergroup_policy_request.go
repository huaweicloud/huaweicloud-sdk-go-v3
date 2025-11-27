package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableClustergroupPolicyRequest Request Object
type DisableClustergroupPolicyRequest struct {

	// 容器舰队id
	Clustergroupid string `json:"clustergroupid"`
}

func (o DisableClustergroupPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableClustergroupPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisableClustergroupPolicyRequest", string(data)}, " ")
}

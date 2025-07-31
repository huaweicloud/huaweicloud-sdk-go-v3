package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AppPolicyInfoList 策略关联主机信息
type AppPolicyInfoList struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 主机id列表
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o AppPolicyInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppPolicyInfoList struct{}"
	}

	return strings.Join([]string{"AppPolicyInfoList", string(data)}, " ")
}

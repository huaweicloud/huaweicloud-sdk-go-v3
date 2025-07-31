package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppWhitelistPolicyHostRequestInfo 策略添加主机
type AddAppWhitelistPolicyHostRequestInfo struct {

	// 主机id列表
	HostIdList []string `json:"host_id_list"`
}

func (o AddAppWhitelistPolicyHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppWhitelistPolicyHostRequestInfo struct{}"
	}

	return strings.Join([]string{"AddAppWhitelistPolicyHostRequestInfo", string(data)}, " ")
}

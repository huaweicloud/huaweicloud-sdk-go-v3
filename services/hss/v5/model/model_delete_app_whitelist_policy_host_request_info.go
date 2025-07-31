package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppWhitelistPolicyHostRequestInfo 策略删除主机
type DeleteAppWhitelistPolicyHostRequestInfo struct {

	// 主机id列表
	HostIdList []string `json:"host_id_list"`
}

func (o DeleteAppWhitelistPolicyHostRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppWhitelistPolicyHostRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteAppWhitelistPolicyHostRequestInfo", string(data)}, " ")
}

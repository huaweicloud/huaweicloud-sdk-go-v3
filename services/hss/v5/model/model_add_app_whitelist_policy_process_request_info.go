package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppWhitelistPolicyProcessRequestInfo 白名单进程列表
type AddAppWhitelistPolicyProcessRequestInfo struct {

	// 白名单进程列表
	ProcessInfoList []AddAppWhitelistPolicyProcessInfo `json:"process_info_list"`
}

func (o AddAppWhitelistPolicyProcessRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppWhitelistPolicyProcessRequestInfo struct{}"
	}

	return strings.Join([]string{"AddAppWhitelistPolicyProcessRequestInfo", string(data)}, " ")
}

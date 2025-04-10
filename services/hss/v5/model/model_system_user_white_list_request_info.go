package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SystemUserWhiteListRequestInfo 系统用户白名单
type SystemUserWhiteListRequestInfo struct {

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 系统用户名列表
	SystemUserNameList *[]string `json:"system_user_name_list,omitempty"`
}

func (o SystemUserWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SystemUserWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"SystemUserWhiteListRequestInfo", string(data)}, " ")
}

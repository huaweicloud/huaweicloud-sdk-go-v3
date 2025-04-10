package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSystemUserWhiteListRequestInfo 系统用户白名单
type AddSystemUserWhiteListRequestInfo struct {

	// 主机ID
	HostId string `json:"host_id"`

	// 系统用户名列表
	SystemUserNameList []string `json:"system_user_name_list"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (o AddSystemUserWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSystemUserWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"AddSystemUserWhiteListRequestInfo", string(data)}, " ")
}

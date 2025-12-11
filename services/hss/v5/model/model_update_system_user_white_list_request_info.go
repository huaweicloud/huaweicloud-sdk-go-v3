package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSystemUserWhiteListRequestInfo 系统用户白名单
type UpdateSystemUserWhiteListRequestInfo struct {

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId string `json:"host_id"`

	// 系统用户名列表
	SystemUserNameList []string `json:"system_user_name_list"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`
}

func (o UpdateSystemUserWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSystemUserWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateSystemUserWhiteListRequestInfo", string(data)}, " ")
}

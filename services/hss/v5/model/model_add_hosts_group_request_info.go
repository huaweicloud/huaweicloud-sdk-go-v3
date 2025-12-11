package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddHostsGroupRequestInfo struct {

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度1-128位
	GroupName string `json:"group_name"`

	// **参数解释**： 服务器ID列表 **取值范围**: 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o AddHostsGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddHostsGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AddHostsGroupRequestInfo", string(data)}, " ")
}

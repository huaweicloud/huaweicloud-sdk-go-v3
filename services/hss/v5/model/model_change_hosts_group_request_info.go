package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChangeHostsGroupRequestInfo struct {

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度0-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 服务器组ID **取值范围**: 字符长度0-64位
	GroupId string `json:"group_id"`

	// **参数解释**: 服务器ID列表 **取值范围**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`
}

func (o ChangeHostsGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeHostsGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"ChangeHostsGroupRequestInfo", string(data)}, " ")
}

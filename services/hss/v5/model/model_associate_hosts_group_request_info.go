package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateHostsGroupRequestInfo struct {

	// **参数解释**: 服务器组ID **约束限制**: 不涉及 **取值范围**: 字符长度0-128位 **默认取值**: 不涉及
	GroupId string `json:"group_id"`

	// **参数解释**： 服务器ID列表 **约束限制**： 不涉及 **取值范围**： 列表长度1-200条 **默认取值**： 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释**: 是否要对全量主机分配到组，如果为true的话，将所有主机分配到组，不需填写host_id_list，如果为false的话，需要填写host_id_list **约束限制**: 不涉及 **取值范围**: - true: 将所有主机分配到组，不需填写host_id_list。 - false: 非全量分配到组，仅对指定的主机列表分配到组，需要填写host_id_list。 **默认取值**: 不涉及
	OperateAll bool `json:"operate_all"`
}

func (o AssociateHostsGroupRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateHostsGroupRequestInfo struct{}"
	}

	return strings.Join([]string{"AssociateHostsGroupRequestInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPortStatusRequestInfo struct {

	// **参数解释**: 端口状态，是否忽略 **约束限制**: 不涉及 **取值范围**: - ignore：忽略 - not_ignore：取消忽略  **默认取值**: 不涉及
	OperateType *string `json:"operate_type,omitempty"`

	// **参数解释**: 主机相关信息列表 **约束限制**: 不涉及 **取值范围**: 最小值0，最大值200 **默认取值**: 不涉及
	DataList *[]PortStatusRequestInfo `json:"data_list,omitempty"`
}

func (o ListPortStatusRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortStatusRequestInfo struct{}"
	}

	return strings.Join([]string{"ListPortStatusRequestInfo", string(data)}, " ")
}

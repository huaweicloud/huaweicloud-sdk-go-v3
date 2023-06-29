package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceGroupRequest Request Object
type ListResourceGroupRequest struct {

	// 资源分组的名称；长度为1-128，只能包含0-9/a-z/A-Z/_/-或汉字；如：ResourceGroup-Test01。
	GroupName *string `json:"group_name,omitempty"`

	// 资源分组的ID，长度为1-128，只能包含0-9/a-z/A-Z；如：rg16063743652226ew93e64p。
	GroupId *string `json:"group_id,omitempty"`

	// 资源分组健康状态，值可为health、unhealth、no_alarm_rule；health表示健康，
	Status *string `json:"status,omitempty"`

	// 分页起始值，类型为integer，默认值为0。
	Start *int32 `json:"start,omitempty"`

	// 单次查询的条数限制，取值范围(0,100]，默认值为100， 用于限制结果数据条数。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListResourceGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceGroupRequest struct{}"
	}

	return strings.Join([]string{"ListResourceGroupRequest", string(data)}, " ")
}

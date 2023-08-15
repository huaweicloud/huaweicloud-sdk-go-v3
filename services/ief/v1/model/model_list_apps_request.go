package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsRequest Request Object
type ListAppsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 应用模板名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 每页显示的条目数量，取值范围1~1000，默认为1000
	Limit *string `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *string `json:"offset,omitempty"`

	// 通过别名过滤，模糊匹配
	Alias *string `json:"alias,omitempty"`

	// public：公共模板，只有管理员才能创建 private：用户创建的应用模板，默认 shared：第三方应用，其他用户共享类型的模板（保留，未实现）
	Visibility *string `json:"visibility,omitempty"`
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}

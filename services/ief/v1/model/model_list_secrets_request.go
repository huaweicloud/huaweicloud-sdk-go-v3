package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListSecretsRequest struct {

	// 铂金版实例ID，专业版实例为空值
	IefInstanceId *string `json:"ief-instance-id,omitempty"`

	// 密钥名称，模糊匹配
	Name *string `json:"name,omitempty"`

	// 每页显示的条目数量，取值范围1~1000，默认值为1000。
	Limit *string `json:"limit,omitempty"`

	// 查询的起始位置，取值范围为非负整数，默认为0
	Offset *string `json:"offset,omitempty"`

	// 排序方式，可根据名称、创建时间、更新时间排序 枚举值： - name - created_at - updated_at 默认升序，如sort=name，降序：sort=name%3Adesc
	Sort *string `json:"sort,omitempty"`
}

func (o ListSecretsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecretsRequest struct{}"
	}

	return strings.Join([]string{"ListSecretsRequest", string(data)}, " ")
}

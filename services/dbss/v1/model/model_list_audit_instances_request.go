package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditInstancesRequest Request Object
type ListAuditInstancesRequest struct {

	// **参数解释**： 分页偏移量，从第一条数据偏移offset条数据后开始查询 **约束限制**： 仅支持大于等于0的整数 **取值范围**： 大于等于0 **默认取值**： 默认值为0
	Offset *string `json:"offset,omitempty"`

	// 查询记录数
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListAuditInstancesRequest", string(data)}, " ")
}

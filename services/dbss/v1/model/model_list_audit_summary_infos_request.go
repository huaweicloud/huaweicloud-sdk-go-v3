package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditSummaryInfosRequest Request Object
type ListAuditSummaryInfosRequest struct {

	// **参数解释**： 分页偏移量，从第一条数据偏移offset条数据后开始查询 **约束限制**： 仅支持大于等于0的整数 **取值范围**： 大于等于0 **默认取值**： 默认值为0
	Offset *string `json:"offset,omitempty"`

	// **参数解释**： 每页查询记录数。 **约束限制**： 仅支持大于0的整数 **取值范围**： 大于0小于等于10000 **默认取值**： 默认值为100
	Limit *string `json:"limit,omitempty"`
}

func (o ListAuditSummaryInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditSummaryInfosRequest struct{}"
	}

	return strings.Join([]string{"ListAuditSummaryInfosRequest", string(data)}, " ")
}

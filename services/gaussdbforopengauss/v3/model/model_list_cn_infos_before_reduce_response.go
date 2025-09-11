package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCnInfosBeforeReduceResponse Response Object
type ListCnInfosBeforeReduceResponse struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// 单次缩容允许最大步长。
	MaxReductionNum *int32 `json:"max_reduction_num,omitempty"`

	// 节点信息列表。
	Nodes          *[]CnInfoBeforeReduce `json:"nodes,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListCnInfosBeforeReduceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCnInfosBeforeReduceResponse struct{}"
	}

	return strings.Join([]string{"ListCnInfosBeforeReduceResponse", string(data)}, " ")
}

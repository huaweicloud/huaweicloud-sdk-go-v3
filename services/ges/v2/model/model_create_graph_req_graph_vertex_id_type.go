package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGraphReqGraphVertexIdType 点的id类型，仅持久化图需要填写。  > 点ID类型确定后，将无法更改，请用户谨慎选择!
type CreateGraphReqGraphVertexIdType struct {

	// id类型，目前支持固定长度fixedLengthString和hash。  - fixedLengthString：固定长度String格式下，实际点ID直接用于内部存储与计算，用户需指定一长度，实际点ID不可超过此长度。长度过大，可能影响查询性能，建议用户根据数据集状态进行设置。  - hash：哈希格式下，内部计算时将实际点ID转换成哈希码进行存储与计算，对实际点ID长度无限制，但是存在极低的概率(约10^(-43))出现点ID碰撞。若用户无法确定点ID的最大长度，建议选择哈希类型。
	IdType string `json:"id_type"`

	// 当id_type取值为fixedLengthString时必填，取值范围：1-128。
	IdLength *int32 `json:"id_length,omitempty"`
}

func (o CreateGraphReqGraphVertexIdType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGraphReqGraphVertexIdType struct{}"
	}

	return strings.Join([]string{"CreateGraphReqGraphVertexIdType", string(data)}, " ")
}

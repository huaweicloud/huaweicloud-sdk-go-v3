package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCn 批量删除CN节点ID信息
type BatchDeleteCn struct {

	// 批量删除CN节点ID
	Instances *[]string `json:"instances,omitempty"`
}

func (o BatchDeleteCn) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCn struct{}"
	}

	return strings.Join([]string{"BatchDeleteCn", string(data)}, " ")
}

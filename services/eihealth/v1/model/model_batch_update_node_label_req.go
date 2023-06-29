package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateNodeLabelReq 待更新的标签集
type BatchUpdateNodeLabelReq struct {

	// 标签列表
	Labels []UpdateNodeLabelReq `json:"labels"`
}

func (o BatchUpdateNodeLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateNodeLabelReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateNodeLabelReq", string(data)}, " ")
}

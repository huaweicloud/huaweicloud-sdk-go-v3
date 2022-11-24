package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 待更新的标签集
type BatchUpadteNodeLabelReq struct {

	// 标签列表
	Labels []UpadteNodeLabelReq `json:"labels"`
}

func (o BatchUpadteNodeLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpadteNodeLabelReq struct{}"
	}

	return strings.Join([]string{"BatchUpadteNodeLabelReq", string(data)}, " ")
}

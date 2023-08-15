package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeLabelReq 标签信息
type UpdateNodeLabelReq struct {

	// 标签名称
	Name string `json:"name"`
}

func (o UpdateNodeLabelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeLabelReq struct{}"
	}

	return strings.Join([]string{"UpdateNodeLabelReq", string(data)}, " ")
}

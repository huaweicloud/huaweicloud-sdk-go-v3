package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点组更新请求体
type EdgeGroupUpdateRequest struct {

	// 边缘节点组描述
	Description *string `json:"description,omitempty"`
}

func (o EdgeGroupUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeGroupUpdateRequest struct{}"
	}

	return strings.Join([]string{"EdgeGroupUpdateRequest", string(data)}, " ")
}

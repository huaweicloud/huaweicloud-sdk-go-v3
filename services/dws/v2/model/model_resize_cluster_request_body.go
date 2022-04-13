package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type ResizeClusterRequestBody struct {
	ScaleOut *ScaleOut `json:"scale_out,omitempty"`
}

func (o ResizeClusterRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeClusterRequestBody struct{}"
	}

	return strings.Join([]string{"ResizeClusterRequestBody", string(data)}, " ")
}

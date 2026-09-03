package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSearchPathFlagRequestBody 设置searchpath开关请求体
type UpdateSearchPathFlagRequestBody struct {

	// 开关标志
	SearchPathFlag bool `json:"search_path_flag"`
}

func (o UpdateSearchPathFlagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSearchPathFlagRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSearchPathFlagRequestBody", string(data)}, " ")
}

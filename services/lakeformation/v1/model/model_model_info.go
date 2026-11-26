package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelInfo model input when grant policy
type ModelInfo struct {

	// 模型名称。只能包含字母、数字、下划线、中划线，且长度为1~256个字符。
	Name string `json:"name"`
}

func (o ModelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelInfo struct{}"
	}

	return strings.Join([]string{"ModelInfo", string(data)}, " ")
}

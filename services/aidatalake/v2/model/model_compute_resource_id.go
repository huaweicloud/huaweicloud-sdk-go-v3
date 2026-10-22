package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputeResourceId **参数解释**：计算资源ID。 **取值范围**：长度为1~36个英文字母、数字和中划线的组合。
type ComputeResourceId struct {
}

func (o ComputeResourceId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeResourceId struct{}"
	}

	return strings.Join([]string{"ComputeResourceId", string(data)}, " ")
}

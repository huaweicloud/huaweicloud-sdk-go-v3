package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelId 模型ID，32~36位的英文、数字、短横组合
type ModelId struct {
}

func (o ModelId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelId struct{}"
	}

	return strings.Join([]string{"ModelId", string(data)}, " ")
}

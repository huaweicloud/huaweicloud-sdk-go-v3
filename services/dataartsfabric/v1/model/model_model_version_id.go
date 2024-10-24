package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelVersionId 模型版本ID，32~36位的英文、数字、短横组合，系统自动生成无法修改，输入不生效。
type ModelVersionId struct {
}

func (o ModelVersionId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelVersionId struct{}"
	}

	return strings.Join([]string{"ModelVersionId", string(data)}, " ")
}

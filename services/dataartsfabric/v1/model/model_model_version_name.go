package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelVersionName 模型版本名称, 只能包含中文、字母、数字、下划线、中划线、点、空格
type ModelVersionName struct {
}

func (o ModelVersionName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelVersionName struct{}"
	}

	return strings.Join([]string{"ModelVersionName", string(data)}, " ")
}

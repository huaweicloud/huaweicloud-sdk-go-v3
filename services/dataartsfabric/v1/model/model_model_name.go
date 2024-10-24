package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelName 一个Model的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
type ModelName struct {
}

func (o ModelName) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelName struct{}"
	}

	return strings.Join([]string{"ModelName", string(data)}, " ")
}

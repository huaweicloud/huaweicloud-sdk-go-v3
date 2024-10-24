package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BaseModelType 类型请从ListBaseModels列举基模型接口响应中获取
type BaseModelType struct {
}

func (o BaseModelType) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseModelType struct{}"
	}

	return strings.Join([]string{"BaseModelType", string(data)}, " ")
}

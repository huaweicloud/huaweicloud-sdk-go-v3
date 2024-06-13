package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAllConfigItemByTypeRequestBody struct {

	// 系统配置，服务自己配置{system、service}
	ConfigType *string `json:"configType,omitempty"`

	// 配置类型集合
	Types *[]string `json:"types,omitempty"`
}

func (o ListAllConfigItemByTypeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllConfigItemByTypeRequestBody struct{}"
	}

	return strings.Join([]string{"ListAllConfigItemByTypeRequestBody", string(data)}, " ")
}

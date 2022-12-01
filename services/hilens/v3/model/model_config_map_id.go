package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigMapId struct {

	// 配置项id
	Id string `json:"id"`
}

func (o ConfigMapId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigMapId struct{}"
	}

	return strings.Join([]string{"ConfigMapId", string(data)}, " ")
}

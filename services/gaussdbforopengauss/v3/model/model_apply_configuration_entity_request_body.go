package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyConfigurationEntityRequestBody struct {

	// **参数解释**: 资源ID列表。可通过“查询可应用实例列表”接口获取。 **约束限制**: 不涉及。
	EntityIds []string `json:"entity_ids"`
}

func (o ApplyConfigurationEntityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyConfigurationEntityRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyConfigurationEntityRequestBody", string(data)}, " ")
}

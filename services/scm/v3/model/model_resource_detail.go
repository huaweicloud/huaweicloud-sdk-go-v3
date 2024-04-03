package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceDetail struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 资源名称。
	ResourceName string `json:"resource_name"`
}

func (o ResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceDetail struct{}"
	}

	return strings.Join([]string{"ResourceDetail", string(data)}, " ")
}

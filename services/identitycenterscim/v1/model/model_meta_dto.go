package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetaDto struct {

	// 资源类型
	ResourceType *string `json:"resourceType,omitempty"`

	// 资源创建时间
	Created *string `json:"created,omitempty"`

	// 资源最后更新时间
	LastModified *string `json:"lastModified,omitempty"`
}

func (o MetaDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaDto struct{}"
	}

	return strings.Join([]string{"MetaDto", string(data)}, " ")
}

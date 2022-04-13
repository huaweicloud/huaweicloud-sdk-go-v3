package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Service struct {
	// 服务名。

	Name string `json:"name"`
	// 服务描述信息。

	Description *string `json:"description,omitempty"`

	Links *Links `json:"links"`
	// 服务ID。

	Id string `json:"id"`
	// 服务类型。

	Type string `json:"type"`
	// 服务是否可用。

	Enabled bool `json:"enabled"`
}

func (o Service) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Service struct{}"
	}

	return strings.Join([]string{"Service", string(data)}, " ")
}

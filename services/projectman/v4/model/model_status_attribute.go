package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatusAttribute struct {

	// 状态属性id
	Id *int32 `json:"id,omitempty"`

	// 状态属性名称
	Name *string `json:"name,omitempty"`
}

func (o StatusAttribute) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatusAttribute struct{}"
	}

	return strings.Join([]string{"StatusAttribute", string(data)}, " ")
}

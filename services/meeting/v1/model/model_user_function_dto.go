package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserFunctionDto struct {

	// 是否开启智能协同白板功能。如果开启，表示该帐号是给智能协同白板使用，占用企业智能协同白板的资源，如果资源不足，则无法开启。 默认值：false。 > 该参数将废弃，请勿使用。
	EnableRoom *bool `json:"enableRoom,omitempty"`
}

func (o UserFunctionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserFunctionDto struct{}"
	}

	return strings.Join([]string{"UserFunctionDto", string(data)}, " ")
}

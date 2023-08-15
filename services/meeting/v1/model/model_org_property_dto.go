package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OrgPropertyDto struct {

	// 配置项key。 开通本地录制功能，参数填写：enableClientRecord。
	PropertyKey *string `json:"propertyKey,omitempty"`

	// 配置项值。 开通本地录制功能，参数填写：true。
	PropertyValue *string `json:"propertyValue,omitempty"`
}

func (o OrgPropertyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OrgPropertyDto struct{}"
	}

	return strings.Join([]string{"OrgPropertyDto", string(data)}, " ")
}

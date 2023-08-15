package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdmetWithCustomPropsResponse Response Object
type ShowAdmetWithCustomPropsResponse struct {

	// 用户已开启的自定义属性集合
	CustomProps *[]CustomProp `json:"custom_props,omitempty"`

	Props          *AdmetPropertyDictWithCustom `json:"props,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowAdmetWithCustomPropsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdmetWithCustomPropsResponse struct{}"
	}

	return strings.Join([]string{"ShowAdmetWithCustomPropsResponse", string(data)}, " ")
}

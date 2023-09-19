package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Property 云手机属性信息。
type Property struct {

	// 云手机id，不超过32个字节。
	PhoneId string `json:"phone_id"`

	// 云手机属性列表，为Json格式字符串。
	Property string `json:"property"`
}

func (o Property) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Property struct{}"
	}

	return strings.Join([]string{"Property", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceTagResp **参数解释**： 资源标签。
type ResourceTagResp struct {

	// **参数解释**： 标签名。 **取值范围**： 长度[1,128]个unicode字符。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 标签值。 **取值范围**： 长度为[0,255]个unicode字符。
	Value *string `json:"value,omitempty"`
}

func (o ResourceTagResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceTagResp struct{}"
	}

	return strings.Join([]string{"ResourceTagResp", string(data)}, " ")
}

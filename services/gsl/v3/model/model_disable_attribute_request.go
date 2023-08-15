package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableAttributeRequest Request Object
type DisableAttributeRequest struct {

	// 自定义属性标识
	AttributeId int64 `json:"attribute_id"`
}

func (o DisableAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableAttributeRequest struct{}"
	}

	return strings.Join([]string{"DisableAttributeRequest", string(data)}, " ")
}

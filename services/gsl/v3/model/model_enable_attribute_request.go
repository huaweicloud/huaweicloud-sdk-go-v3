package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableAttributeRequest Request Object
type EnableAttributeRequest struct {

	// 自定义属性标识
	AttributeId int64 `json:"attribute_id"`
}

func (o EnableAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableAttributeRequest struct{}"
	}

	return strings.Join([]string{"EnableAttributeRequest", string(data)}, " ")
}

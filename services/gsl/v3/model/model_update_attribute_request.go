package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAttributeRequest Request Object
type UpdateAttributeRequest struct {

	// 自定义属性标识
	AttributeId int64 `json:"attribute_id"`

	Body *AddOrModifyAttributeReq `json:"body,omitempty"`
}

func (o UpdateAttributeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAttributeRequest struct{}"
	}

	return strings.Join([]string{"UpdateAttributeRequest", string(data)}, " ")
}

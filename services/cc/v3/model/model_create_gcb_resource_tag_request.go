package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGcbResourceTagRequest Request Object
type CreateGcbResourceTagRequest struct {

	// 资源唯一标识符。
	ResourceId string `json:"resource_id"`

	Body *CreateGcbTagRequestBody `json:"body,omitempty"`
}

func (o CreateGcbResourceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGcbResourceTagRequest struct{}"
	}

	return strings.Join([]string{"CreateGcbResourceTagRequest", string(data)}, " ")
}

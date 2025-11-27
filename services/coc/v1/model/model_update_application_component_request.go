package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationComponentRequest Request Object
type UpdateApplicationComponentRequest struct {

	// 分组ID。
	Id string `json:"id"`

	Body *ComponentUpdateRequest `json:"body,omitempty"`
}

func (o UpdateApplicationComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationComponentRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationComponentRequest", string(data)}, " ")
}

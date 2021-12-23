package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddProtectedInstanceTagsRequest struct {
	// 保护实例的ID。

	ProtectedInstanceId string `json:"protected_instance_id"`

	Body *ProtectedInstanceAddTagsRequestBody `json:"body,omitempty"`
}

func (o AddProtectedInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceTagsRequest", string(data)}, " ")
}

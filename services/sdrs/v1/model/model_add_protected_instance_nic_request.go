package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddProtectedInstanceNicRequest struct {

	// 保护实例的ID。
	ProtectedInstanceId string `json:"protected_instance_id" xml:"protected_instance_id"`

	Body *ProtectedInstanceAddNicRequestBody `json:"body,omitempty" xml:"body"`
}

func (o AddProtectedInstanceNicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddProtectedInstanceNicRequest struct{}"
	}

	return strings.Join([]string{"AddProtectedInstanceNicRequest", string(data)}, " ")
}

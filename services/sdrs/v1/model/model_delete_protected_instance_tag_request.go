package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteProtectedInstanceTagRequest struct {

	// 保护实例的ID。
	ProtectedInstanceId string `json:"protected_instance_id" xml:"protected_instance_id"`

	// 标签key。
	Key string `json:"key" xml:"key"`
}

func (o DeleteProtectedInstanceTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteProtectedInstanceTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteProtectedInstanceTagRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProtectedInstanceTagsRequest struct {

	// 保护实例的ID。
	ProtectedInstanceId string `json:"protected_instance_id" xml:"protected_instance_id"`
}

func (o ListProtectedInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProtectedInstanceTagsRequest", string(data)}, " ")
}

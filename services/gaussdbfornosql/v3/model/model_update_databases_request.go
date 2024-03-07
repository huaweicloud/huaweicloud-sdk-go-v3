package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDatabasesRequest Request Object
type UpdateDatabasesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateDatabasesRequestBody `json:"body,omitempty"`
}

func (o UpdateDatabasesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDatabasesRequest struct{}"
	}

	return strings.Join([]string{"UpdateDatabasesRequest", string(data)}, " ")
}

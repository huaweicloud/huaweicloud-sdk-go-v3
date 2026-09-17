package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAllSessionsRequest Request Object
type DeleteAllSessionsRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o DeleteAllSessionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAllSessionsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAllSessionsRequest", string(data)}, " ")
}

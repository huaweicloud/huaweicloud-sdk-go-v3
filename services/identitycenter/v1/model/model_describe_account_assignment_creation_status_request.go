package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeAccountAssignmentCreationStatusRequest Request Object
type DescribeAccountAssignmentCreationStatusRequest struct {
	InstanceId string `json:"instance_id"`

	// 请求的唯一标识
	RequestId string `json:"request_id"`
}

func (o DescribeAccountAssignmentCreationStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeAccountAssignmentCreationStatusRequest struct{}"
	}

	return strings.Join([]string{"DescribeAccountAssignmentCreationStatusRequest", string(data)}, " ")
}

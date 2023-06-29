package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeAccountAssignmentDeletionStatusRequest Request Object
type DescribeAccountAssignmentDeletionStatusRequest struct {
	InstanceId string `json:"instance_id"`

	// 请求唯一标识
	RequestId string `json:"request_id"`
}

func (o DescribeAccountAssignmentDeletionStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribeAccountAssignmentDeletionStatusRequest struct{}"
	}

	return strings.Join([]string{"DescribeAccountAssignmentDeletionStatusRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribeAccountAssignmentCreationStatusRequest Request Object
type DescribeAccountAssignmentCreationStatusRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// IAM Identity Center实例的全局唯一标识符（ID）
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

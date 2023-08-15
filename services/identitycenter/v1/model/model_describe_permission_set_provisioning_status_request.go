package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribePermissionSetProvisioningStatusRequest Request Object
type DescribePermissionSetProvisioningStatusRequest struct {
	InstanceId string `json:"instance_id"`

	RequestId string `json:"request_id"`
}

func (o DescribePermissionSetProvisioningStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribePermissionSetProvisioningStatusRequest struct{}"
	}

	return strings.Join([]string{"DescribePermissionSetProvisioningStatusRequest", string(data)}, " ")
}

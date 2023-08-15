package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribePermissionSetRequest Request Object
type DescribePermissionSetRequest struct {
	InstanceId string `json:"instance_id"`

	PermissionSetId string `json:"permission_set_id"`
}

func (o DescribePermissionSetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribePermissionSetRequest struct{}"
	}

	return strings.Join([]string{"DescribePermissionSetRequest", string(data)}, " ")
}

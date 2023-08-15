package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DescribePermissionSetResponse Response Object
type DescribePermissionSetResponse struct {
	PermissionSet  *PermissionSetDto `json:"permission_set,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DescribePermissionSetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DescribePermissionSetResponse struct{}"
	}

	return strings.Join([]string{"DescribePermissionSetResponse", string(data)}, " ")
}

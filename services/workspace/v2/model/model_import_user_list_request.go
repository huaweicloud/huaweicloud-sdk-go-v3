package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportUserListRequest Request Object
type ImportUserListRequest struct {

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网id。
	SubnetId *string `json:"subnet_id,omitempty"`

	Body *ImportUserListRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportUserListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportUserListRequest struct{}"
	}

	return strings.Join([]string{"ImportUserListRequest", string(data)}, " ")
}

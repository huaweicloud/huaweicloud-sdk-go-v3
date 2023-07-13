package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IamGroupDto struct {

	// IAM用户组id
	Id *string `json:"id,omitempty"`

	// IAM用户组名
	Name *string `json:"name,omitempty"`
}

func (o IamGroupDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamGroupDto struct{}"
	}

	return strings.Join([]string{"IamGroupDto", string(data)}, " ")
}

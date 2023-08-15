package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IamUserDto struct {

	// IAM用户id
	Id *string `json:"id,omitempty"`

	// IAM用户名
	Name *string `json:"name,omitempty"`
}

func (o IamUserDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamUserDto struct{}"
	}

	return strings.Join([]string{"IamUserDto", string(data)}, " ")
}

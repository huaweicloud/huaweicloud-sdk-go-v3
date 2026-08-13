package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GenericResourceGroupRequestBody struct {
	GroupName string `json:"group_name"`

	Description *string `json:"description,omitempty"`
}

func (o GenericResourceGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenericResourceGroupRequestBody struct{}"
	}

	return strings.Join([]string{"GenericResourceGroupRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AddGroupRequestBody struct {

	// 群组名称
	Name *string `json:"name,omitempty"`
}

func (o AddGroupRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddGroupRequestBody struct{}"
	}

	return strings.Join([]string{"AddGroupRequestBody", string(data)}, " ")
}

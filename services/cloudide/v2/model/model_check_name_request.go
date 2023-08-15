package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckNameRequest Request Object
type CheckNameRequest struct {

	// 实例名。 可以输入中文、数字、字母、下划线、点、破折号。长度介于3-100之间
	DisplayName string `json:"display_name"`
}

func (o CheckNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckNameRequest struct{}"
	}

	return strings.Join([]string{"CheckNameRequest", string(data)}, " ")
}

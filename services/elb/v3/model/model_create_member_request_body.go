package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMemberRequestBody This is a auto create Body Object
type CreateMemberRequestBody struct {
	Member *CreateMemberOption `json:"member"`
}

func (o CreateMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMemberRequestBody", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// This is a auto create Body Object
type CreateMemberRequestBody struct {
	Member *CreateMemberReq `json:"member"`
}

func (o CreateMemberRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMemberRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMemberRequestBody", string(data)}, " ")
}

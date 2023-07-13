package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateUserCenterAuthorizationV2Req struct {

	// 是否同意 0不同意 1同意
	IsAuthorized *int32 `json:"is_authorized,omitempty"`

	// 机密信息内容
	AuthorizationContent *string `json:"authorization_content,omitempty"`
}

func (o CreateUserCenterAuthorizationV2Req) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserCenterAuthorizationV2Req struct{}"
	}

	return strings.Join([]string{"CreateUserCenterAuthorizationV2Req", string(data)}, " ")
}

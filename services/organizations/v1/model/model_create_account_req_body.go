package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccount 操作的请求体。
type CreateAccountReqBody struct {

	// 帐号名称。
	Name string `json:"name"`

	// 要绑定到新创建的帐号的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreateAccountReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountReqBody struct{}"
	}

	return strings.Join([]string{"CreateAccountReqBody", string(data)}, " ")
}

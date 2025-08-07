package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceAccountReqBody CreateAccount 操作的请求体。
type CreateResourceAccountReqBody struct {

	// 帐号名称。
	Name string `json:"name"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 描述信息。
	Description *string `json:"description,omitempty"`

	// 要绑定到新创建的帐号的标签列表。
	Tags *[]TagDto `json:"tags,omitempty"`
}

func (o CreateResourceAccountReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceAccountReqBody struct{}"
	}

	return strings.Join([]string{"CreateResourceAccountReqBody", string(data)}, " ")
}

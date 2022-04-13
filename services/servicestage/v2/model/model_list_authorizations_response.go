package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAuthorizationsResponse struct {
	// 授权列表。

	Authorizations *[]AuthorizationVo `json:"authorizations,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizationsResponse", string(data)}, " ")
}

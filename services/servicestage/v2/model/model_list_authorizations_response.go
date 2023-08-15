package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuthorizationsResponse Response Object
type ListAuthorizationsResponse struct {

	// 授权列表。
	Authorizations *[]AuthorizationVo `json:"authorizations,omitempty"`

	// 仓库授权数量。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuthorizationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizationsResponse", string(data)}, " ")
}

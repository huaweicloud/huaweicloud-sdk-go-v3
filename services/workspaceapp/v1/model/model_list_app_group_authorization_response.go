package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppGroupAuthorizationResponse Response Object
type ListAppGroupAuthorizationResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 授权信息。
	Authorizations *[]Authorization `json:"authorizations,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAppGroupAuthorizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppGroupAuthorizationResponse struct{}"
	}

	return strings.Join([]string{"ListAppGroupAuthorizationResponse", string(data)}, " ")
}

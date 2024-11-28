package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopPoolAuthorizedObjectsResponse Response Object
type ListDesktopPoolAuthorizedObjectsResponse struct {

	// 授权对象
	Objects *[]AuthorizedObjects `json:"objects,omitempty"`

	// 满足条件的用户、用户组总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopPoolAuthorizedObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopPoolAuthorizedObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopPoolAuthorizedObjectsResponse", string(data)}, " ")
}

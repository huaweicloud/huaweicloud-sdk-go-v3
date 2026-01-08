package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckEdgeSiteResourcesResponse Response Object
type CheckEdgeSiteResourcesResponse struct {

	// 配额是否足够1：足够 0 ：不足。
	IsEnough *int32 `json:"is_enough,omitempty"`

	// 资源剩余数量信息。
	ResourceRemainder *[]ResourceRemainderData `json:"resource_remainder,omitempty"`
	HttpStatusCode    int                      `json:"-"`
}

func (o CheckEdgeSiteResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckEdgeSiteResourcesResponse struct{}"
	}

	return strings.Join([]string{"CheckEdgeSiteResourcesResponse", string(data)}, " ")
}

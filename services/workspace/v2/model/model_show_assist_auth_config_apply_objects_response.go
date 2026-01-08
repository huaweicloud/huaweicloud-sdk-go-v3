package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssistAuthConfigApplyObjectsResponse Response Object
type ShowAssistAuthConfigApplyObjectsResponse struct {

	// 应用对象。
	Objects *[]ApplyObjects `json:"objects,omitempty"`

	// 满足条件的用户、用户组总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAssistAuthConfigApplyObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssistAuthConfigApplyObjectsResponse struct{}"
	}

	return strings.Join([]string{"ShowAssistAuthConfigApplyObjectsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDesktopPoolAuthorizedObjectsRequest Request Object
type CreateDesktopPoolAuthorizedObjectsRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	Body *CreateDesktopPoolAuthorizedObjectsRequestBody `json:"body,omitempty"`
}

func (o CreateDesktopPoolAuthorizedObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDesktopPoolAuthorizedObjectsRequest struct{}"
	}

	return strings.Join([]string{"CreateDesktopPoolAuthorizedObjectsRequest", string(data)}, " ")
}

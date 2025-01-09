package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAppAuthorizationsRequest Request Object
type UpdateAppAuthorizationsRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	Body *AssignAppAuthorizationsReq `json:"body,omitempty"`
}

func (o UpdateAppAuthorizationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppAuthorizationsRequest struct{}"
	}

	return strings.Join([]string{"UpdateAppAuthorizationsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUploadedAppRequest Request Object
type UpdateUploadedAppRequest struct {

	// 应用ID。
	AppId string `json:"app_id"`

	Body *UpdateAppReq `json:"body,omitempty"`
}

func (o UpdateUploadedAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUploadedAppRequest struct{}"
	}

	return strings.Join([]string{"UpdateUploadedAppRequest", string(data)}, " ")
}

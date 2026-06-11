package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAppIconRawRequest Request Object
type UploadAppIconRawRequest struct {

	// 应用组ID。
	AppGroupId string `json:"app_group_id"`

	// 应用ID。
	AppId string `json:"app_id"`

	Body *UpdateRawIconReq `json:"body,omitempty"`
}

func (o UploadAppIconRawRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAppIconRawRequest struct{}"
	}

	return strings.Join([]string{"UploadAppIconRawRequest", string(data)}, " ")
}

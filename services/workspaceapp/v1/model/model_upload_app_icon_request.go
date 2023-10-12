package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UploadAppIconRequest Request Object
type UploadAppIconRequest struct {

	// 应用组ID
	AppGroupId string `json:"app_group_id"`

	// 应用ID
	AppId string `json:"app_id"`

	Body *UploadAppIconRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o UploadAppIconRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadAppIconRequest struct{}"
	}

	return strings.Join([]string{"UploadAppIconRequest", string(data)}, " ")
}

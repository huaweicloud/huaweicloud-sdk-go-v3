package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadAppVersionRequest Request Object
type DownloadAppVersionRequest struct {

	// 应用ID
	AppId string `json:"app_id"`

	// 应用版本
	Version string `json:"version"`
}

func (o DownloadAppVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAppVersionRequest struct{}"
	}

	return strings.Join([]string{"DownloadAppVersionRequest", string(data)}, " ")
}

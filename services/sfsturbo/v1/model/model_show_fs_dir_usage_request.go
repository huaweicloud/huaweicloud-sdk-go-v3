package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFsDirUsageRequest Request Object
type ShowFsDirUsageRequest struct {

	// MIME类型, application/json
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 文件系统内合法的目录全路径
	Path string `json:"path"`
}

func (o ShowFsDirUsageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsDirUsageRequest struct{}"
	}

	return strings.Join([]string{"ShowFsDirUsageRequest", string(data)}, " ")
}

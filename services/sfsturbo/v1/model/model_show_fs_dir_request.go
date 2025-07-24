package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFsDirRequest Request Object
type ShowFsDirRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 需要查询的目录的全路径
	Path string `json:"path"`
}

func (o ShowFsDirRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsDirRequest struct{}"
	}

	return strings.Join([]string{"ShowFsDirRequest", string(data)}, " ")
}

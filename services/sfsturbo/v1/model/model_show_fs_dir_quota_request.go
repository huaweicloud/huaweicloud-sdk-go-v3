package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFsDirQuotaRequest Request Object
type ShowFsDirQuotaRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 合法的已存在的目录的全路径
	Path string `json:"path"`
}

func (o ShowFsDirQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFsDirQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowFsDirQuotaRequest", string(data)}, " ")
}

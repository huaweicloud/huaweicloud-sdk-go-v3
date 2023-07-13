package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShareRequest Request Object
type ShowShareRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`
}

func (o ShowShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShareRequest struct{}"
	}

	return strings.Join([]string{"ShowShareRequest", string(data)}, " ")
}

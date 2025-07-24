package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClientIpInfoRequest Request Object
type ShowClientIpInfoRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *ShowClientIpInfoRequestBody `json:"body,omitempty"`
}

func (o ShowClientIpInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClientIpInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowClientIpInfoRequest", string(data)}, " ")
}

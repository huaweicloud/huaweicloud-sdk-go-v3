package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackendTargetInfoRequest Request Object
type ShowBackendTargetInfoRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 绑定关系id
	TargetId string `json:"target_id"`
}

func (o ShowBackendTargetInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackendTargetInfoRequest struct{}"
	}

	return strings.Join([]string{"ShowBackendTargetInfoRequest", string(data)}, " ")
}

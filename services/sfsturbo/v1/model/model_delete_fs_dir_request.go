package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFsDirRequest Request Object
type DeleteFsDirRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	Body *DeleteFsDirRequestBody `json:"body,omitempty"`
}

func (o DeleteFsDirRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFsDirRequest struct{}"
	}

	return strings.Join([]string{"DeleteFsDirRequest", string(data)}, " ")
}

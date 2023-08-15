package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteShareRequest Request Object
type DeleteShareRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`
}

func (o DeleteShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareRequest struct{}"
	}

	return strings.Join([]string{"DeleteShareRequest", string(data)}, " ")
}

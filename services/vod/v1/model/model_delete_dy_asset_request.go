package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDyAssetRequest Request Object
type DeleteDyAssetRequest struct {

	// obs桶名称
	Bucket string `json:"bucket"`

	// obs文件路径，不包含桶名
	Object []string `json:"object"`

	// 回调地址
	CallbackUrl *string `json:"callback_url,omitempty"`

	// 用户透传信息
	SessionContext *string `json:"session_context,omitempty"`
}

func (o DeleteDyAssetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDyAssetRequest struct{}"
	}

	return strings.Join([]string{"DeleteDyAssetRequest", string(data)}, " ")
}

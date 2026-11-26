package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetExtendMessageResponse Response Object
type ShowAssetExtendMessageResponse struct {

	// 资产在README中的扩展内容。
	Body           map[string]interface{} `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowAssetExtendMessageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetExtendMessageResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetExtendMessageResponse", string(data)}, " ")
}

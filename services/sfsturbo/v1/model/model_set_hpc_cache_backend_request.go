package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetHpcCacheBackendRequest Request Object
type SetHpcCacheBackendRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	Body *ReqConfigHpcCacheBackend `json:"body,omitempty"`
}

func (o SetHpcCacheBackendRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetHpcCacheBackendRequest struct{}"
	}

	return strings.Join([]string{"SetHpcCacheBackendRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchChangeMaintainServerReq 标记指定的服务器维护状态请求。
type BatchChangeMaintainServerReq struct {

	// 批量请求的服务器ID列表，一次请求数量区间 [1, 20]。
	Items []string `json:"items"`

	// 服务器维护状态标识： * `true` - 添加标记 * `false` - 移除标记
	MaintainStatus bool `json:"maintain_status"`
}

func (o BatchChangeMaintainServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchChangeMaintainServerReq struct{}"
	}

	return strings.Join([]string{"BatchChangeMaintainServerReq", string(data)}, " ")
}

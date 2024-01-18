package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateServerReq 更新服务器请求。
type UpdateServerReq struct {

	// 服务器名称，名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成。 2. 长度范围1~64个字符。
	Name *string `json:"name,omitempty"`

	// 服务器描述。
	Description *string `json:"description,omitempty"`

	// 服务器维护状态标识： * `true` - 添加标记 * `false` - 移除标记
	MaintainStatus *bool `json:"maintain_status,omitempty"`
}

func (o UpdateServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServerReq struct{}"
	}

	return strings.Join([]string{"UpdateServerReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateImageServerReq 更新镜像实例请求。
type UpdateImageServerReq struct {

	// 镜像实例名称，名称需满足如下规则: 1. 由中文，英文大小写，数字，_-组成，不能有空格。 2. 长度范围1~64个字符。
	Name *string `json:"name,omitempty"`

	// 服务器组描述。
	Description *string `json:"description,omitempty"`
}

func (o UpdateImageServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImageServerReq struct{}"
	}

	return strings.Join([]string{"UpdateImageServerReq", string(data)}, " ")
}

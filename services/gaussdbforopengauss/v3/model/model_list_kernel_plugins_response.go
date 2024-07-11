package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListKernelPluginsResponse Response Object
type ListKernelPluginsResponse struct {

	// 插件列表
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListKernelPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKernelPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListKernelPluginsResponse", string(data)}, " ")
}

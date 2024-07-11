package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSupportKernelPluginsResponse Response Object
type ListSupportKernelPluginsResponse struct {

	// 插件列表
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportKernelPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportKernelPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListSupportKernelPluginsResponse", string(data)}, " ")
}

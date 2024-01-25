package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginsResponse Response Object
type ShowPluginsResponse struct {

	// 插件总数。
	Total *int32 `json:"total,omitempty"`

	// 插件列表。
	Data           *[]WasmPlugin `json:"data,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginsResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginsResponse", string(data)}, " ")
}

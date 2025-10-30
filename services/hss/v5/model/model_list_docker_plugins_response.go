package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDockerPluginsResponse Response Object
type ListDockerPluginsResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// data list
	DataList       *[]PluginResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDockerPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDockerPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListDockerPluginsResponse", string(data)}, " ")
}

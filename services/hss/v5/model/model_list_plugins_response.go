package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginsResponse Response Object
type ListPluginsResponse struct {

	// **参数解释**： 插件信息列表 **取值范围**: 不涉及
	DataList *[]PluginInfo `json:"data_list,omitempty"`

	// **参数解释**： 插件总数 **取值范围**: 不涉及
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListPluginsResponse", string(data)}, " ")
}

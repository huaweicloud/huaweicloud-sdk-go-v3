package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginStatusResponse Response Object
type ListPluginStatusResponse struct {

	// **参数解释**： 总数 **取值范围**: 不涉及
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 插件状态信息列表 **取值范围**: 不涉及
	DataList       *[]PluginStatusInfo `json:"data_list,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListPluginStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginStatusResponse struct{}"
	}

	return strings.Join([]string{"ListPluginStatusResponse", string(data)}, " ")
}

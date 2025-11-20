package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginInfoResponse Response Object
type ListPluginInfoResponse struct {

	// **参数解释**： 插件详细信息列表 **取值范围**: 不涉及
	DataList *[]PluginDetailInfo `json:"data_list,omitempty"`

	// **参数解释**： 总数 **取值范围**: 不涉及
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPluginInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginInfoResponse struct{}"
	}

	return strings.Join([]string{"ListPluginInfoResponse", string(data)}, " ")
}

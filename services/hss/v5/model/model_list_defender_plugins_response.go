package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDefenderPluginsResponse Response Object
type ListDefenderPluginsResponse struct {

	// **参数解释**： 总记录数 **取值范围**： 最小值0，最大值2000000
	TotalNum *int32 `json:"total_num,omitempty"`

	// **参数解释**： 插件列表 **取值范围**： 0~200条
	DataList       *[]ListDefenderPluginsResponseBodyData `json:"data_list,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o ListDefenderPluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDefenderPluginsResponse struct{}"
	}

	return strings.Join([]string{"ListDefenderPluginsResponse", string(data)}, " ")
}

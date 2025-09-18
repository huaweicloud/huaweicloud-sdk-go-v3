package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStagePluginsResponse Response Object
type ListStagePluginsResponse struct {

	// **参数解释**： 结果集。 **取值范围**： 不涉及。
	FullStagePluginsItemList *[]FullStagePluginsRelationVoFullStagePluginsItemList `json:"full_stage_plugins_item_list,omitempty"`
	HttpStatusCode           int                                                   `json:"-"`
}

func (o ListStagePluginsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStagePluginsResponse struct{}"
	}

	return strings.Join([]string{"ListStagePluginsResponse", string(data)}, " ")
}

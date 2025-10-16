package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginInfoListResponse Response Object
type ListPluginInfoListResponse struct {

	// **参数解释**: 插件数量。 **取值范围**: 不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**: 插件详细信息。
	Plugins        *[]CustomerPluginInfoResult `json:"plugins,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListPluginInfoListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginInfoListResponse struct{}"
	}

	return strings.Join([]string{"ListPluginInfoListResponse", string(data)}, " ")
}

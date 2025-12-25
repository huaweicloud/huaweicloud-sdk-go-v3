package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLtsLogsResponse Response Object
type ListLtsLogsResponse struct {

	// **参数解释**： 日志开启状态。 **取值范围**： - OPEN：开启。 - CLOSE：关闭。 - DELETE：集群已删除。
	AccessStatus *string `json:"access_status,omitempty"`

	// **参数解释**： LTS日志列表。 **取值范围**： 不涉及。
	LtsAccessList *[]LtslogInfo `json:"lts_access_list,omitempty"`

	// **参数解释**： 总数量。 **取值范围**： 大于等于0。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLtsLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLtsLogsResponse struct{}"
	}

	return strings.Join([]string{"ListLtsLogsResponse", string(data)}, " ")
}

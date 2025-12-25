package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookStatisticsRequest Request Object
type ShowPlaybookStatisticsRequest struct {

	// **参数解释：** 内容类型 - application/json    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json  **默认取值：** application/json
	ContentType string `json:"content-type"`

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`
}

func (o ShowPlaybookStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookStatisticsRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookMonitorsRequest Request Object
type ShowPlaybookMonitorsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// ID of workspace
	WorkspaceId string `json:"workspace_id"`

	// ID of playbook
	PlaybookId string `json:"playbook_id"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 统计剧本版本类型（ALL:全部，VALID:有效的，DELETED:已删除）
	VersionQueryType string `json:"version_query_type"`

	// 结束时间
	EndTime string `json:"end_time"`
}

func (o ShowPlaybookMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ShowPlaybookMonitorsRequest", string(data)}, " ")
}

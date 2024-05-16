package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityAssignedQueuesResponse Response Object
type ListSecurityAssignedQueuesResponse struct {

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 队列资源信息。
	QueueSources   *[]QueueSrcAssignEntity `json:"queue_sources,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListSecurityAssignedQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityAssignedQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityAssignedQueuesResponse", string(data)}, " ")
}

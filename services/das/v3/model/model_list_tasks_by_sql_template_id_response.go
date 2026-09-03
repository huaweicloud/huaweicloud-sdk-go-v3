package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTasksBySqlTemplateIdResponse Response Object
type ListTasksBySqlTemplateIdResponse struct {

	// SQL洞察任务明细
	Tasks *[]SqlParseTask `json:"tasks,omitempty"`

	// 总数
	Total          *int64 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTasksBySqlTemplateIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTasksBySqlTemplateIdResponse struct{}"
	}

	return strings.Join([]string{"ListTasksBySqlTemplateIdResponse", string(data)}, " ")
}

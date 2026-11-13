package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplicationsResponse Response Object
type ListApplicationsResponse struct {

	// **参数解释**： 返回的应用列表信息。 **取值范围**： 不涉及。
	ApplicationList *[]Application `json:"application_list,omitempty"`

	// **参数解释**： 总数。 **取值范围**： 不涉及。
	Total *int64 `json:"total,omitempty"`

	XTraceId       *string `json:"X-Trace-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListApplicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplicationsResponse struct{}"
	}

	return strings.Join([]string{"ListApplicationsResponse", string(data)}, " ")
}

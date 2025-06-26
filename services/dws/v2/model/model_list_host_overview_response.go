package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostOverviewResponse Response Object
type ListHostOverviewResponse struct {

	// **参数解释**： 查询主机概览响应。 **取值范围**： 不涉及。
	Body           *[]HostOverviewResponse `json:"body,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListHostOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostOverviewResponse struct{}"
	}

	return strings.Join([]string{"ListHostOverviewResponse", string(data)}, " ")
}

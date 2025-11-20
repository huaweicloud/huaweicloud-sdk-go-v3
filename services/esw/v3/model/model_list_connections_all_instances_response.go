package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsAllInstancesResponse Response Object
type ListConnectionsAllInstancesResponse struct {

	// - 参数解释：请求的唯一标识。 - 约束限制：UUID格式。 - 取值范围：不涉及。 - 默认取值：不涉及。
	RequestId *string `json:"request_id,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// - 参数解释：查询二层连接列表的响应体。 - 约束限制：不涉及。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Connections    *[]Connection `json:"connections,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListConnectionsAllInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsAllInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListConnectionsAllInstancesResponse", string(data)}, " ")
}

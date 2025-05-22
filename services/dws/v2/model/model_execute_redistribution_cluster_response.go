package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteRedistributionClusterResponse Response Object
type ExecuteRedistributionClusterResponse struct {

	// **参数解释**： 请求成功时的空白响应。 **取值范围**： 不涉及。
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ExecuteRedistributionClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteRedistributionClusterResponse struct{}"
	}

	return strings.Join([]string{"ExecuteRedistributionClusterResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDwsClusterResponse Response Object
type DeleteDwsClusterResponse struct {

	// **参数解释**： 错误码。 **取值范围**： 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **取值范围**： 不涉及。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// **参数解释**： 任务ID，可用于查询任务进度信息。 **取值范围**： null表示该接口不返回异步任务信息，其它值为异步任务的ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDwsClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDwsClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteDwsClusterResponse", string(data)}, " ")
}

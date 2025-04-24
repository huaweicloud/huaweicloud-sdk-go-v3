package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTaurusDbNodeProcessesRequestBody 终止TaurusDB节点用户会话线程请求
type DeleteTaurusDbNodeProcessesRequestBody struct {

	// **参数解释**：  待终止用户会话线程ID列表。  通过查询节点用户会话线程接口，或show processlist命令获取。
	Processes []int64 `json:"processes"`
}

func (o DeleteTaurusDbNodeProcessesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTaurusDbNodeProcessesRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteTaurusDbNodeProcessesRequestBody", string(data)}, " ")
}

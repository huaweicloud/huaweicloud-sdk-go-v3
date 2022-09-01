package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量暂停任务请求体
type BatchPauseJobReq struct {

	// 不能包含空对象。 集合中的job_id取值严格匹配UUID规则。
	Jobs []PauseInfo `json:"jobs" xml:"jobs"`
}

func (o BatchPauseJobReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPauseJobReq struct{}"
	}

	return strings.Join([]string{"BatchPauseJobReq", string(data)}, " ")
}

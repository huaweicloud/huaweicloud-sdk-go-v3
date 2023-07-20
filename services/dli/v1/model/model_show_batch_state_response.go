package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchStateResponse Response Object
type ShowBatchStateResponse struct {

	// 批处理作业的ID，采用UUID（通用唯一识别码）格式。
	Id *string `json:"id,omitempty"`

	// 批处理作业的状态。starting：正在启动；running：正在执行任务；dead：session已退出；success：session停止成功；recovering：正在恢复。
	State          *string `json:"state,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBatchStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchStateResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchStateResponse", string(data)}, " ")
}

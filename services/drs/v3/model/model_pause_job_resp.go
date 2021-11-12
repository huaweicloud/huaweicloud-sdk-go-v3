package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 暂停任务返回体
type PauseJobResp struct {
	// 任务id

	Id string `json:"id"`
	// 暂停结果

	Status string `json:"status"`
	// 错误码

	ErrorCode *string `json:"error_code,omitempty"`
	// 错误信息

	ErrorMsg *string `json:"error_msg,omitempty"`
}

func (o PauseJobResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PauseJobResp struct{}"
	}

	return strings.Join([]string{"PauseJobResp", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExpandGraphResponse struct {
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。

	ErrorMessage *string `json:"errorMessage,omitempty"`
	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。

	ErrorCode *string `json:"errorCode,omitempty"`
	// 扩副本任务ID。请求失败时字段为空。 >可以查询jobId查看任务执行状态、获取返回结果

	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExpandGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpandGraphResponse struct{}"
	}

	return strings.Join([]string{"ExpandGraphResponse", string(data)}, " ")
}

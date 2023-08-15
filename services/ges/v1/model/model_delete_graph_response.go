package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGraphResponse Response Object
type DeleteGraphResponse struct {

	// 删除图任务ID。请求失败时字段为空。 >可以查询jobId查看任务执行状态、获取返回结果
	JobId *string `json:"jobId,omitempty"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误信息。
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// 系统提示信息，执行成功时，字段可能为空。执行失败时，用于显示错误码。
	ErrorCode      *string `json:"errorCode,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteGraphResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGraphResponse struct{}"
	}

	return strings.Join([]string{"DeleteGraphResponse", string(data)}, " ")
}

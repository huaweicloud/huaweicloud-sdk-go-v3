package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunQueryDocumentModerationJobResponse Response Object
type RunQueryDocumentModerationJobResponse struct {

	// 作业id
	JobId *string `json:"job_id,omitempty"`

	// 作业状态，可取值有： running: 正在运行 succeeded: 运行成功 failed: 运行失败
	Status *string `json:"status,omitempty"`

	Result *DocumentQueryResponseResult `json:"result,omitempty"`

	RequestParams *DocumentQueryResponseRequestParams `json:"request_params,omitempty"`

	// 作业创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 作业更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 本次请求的唯⼀标识，⽤于问题排查，建议保存。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunQueryDocumentModerationJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryDocumentModerationJobResponse struct{}"
	}

	return strings.Join([]string{"RunQueryDocumentModerationJobResponse", string(data)}, " ")
}

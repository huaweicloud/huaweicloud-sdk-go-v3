package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMemberHealthCheckJobRequest Request Object
type ShowMemberHealthCheckJobRequest struct {

	// 参数解释：创建后端服务器检测任务的接口返回的job_id。
	JobId string `json:"job_id"`
}

func (o ShowMemberHealthCheckJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMemberHealthCheckJobRequest struct{}"
	}

	return strings.Join([]string{"ShowMemberHealthCheckJobRequest", string(data)}, " ")
}

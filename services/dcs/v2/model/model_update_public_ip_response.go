package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicIpResponse Response Object
type UpdatePublicIpResponse struct {

	// 请求提交任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicIpResponse struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {
	Job *MainJob `json:"job,omitempty"`

	// 请求ID。 注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

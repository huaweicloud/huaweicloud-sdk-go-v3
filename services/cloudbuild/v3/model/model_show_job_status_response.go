package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobStatusResponse struct {
	// 任务运行结果

	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowJobStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowJobStatusResponse", string(data)}, " ")
}

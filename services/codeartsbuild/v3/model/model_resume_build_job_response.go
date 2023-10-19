package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResumeBuildJobResponse Response Object
type ResumeBuildJobResponse struct {

	// 返回结果状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResumeBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumeBuildJobResponse struct{}"
	}

	return strings.Join([]string{"ResumeBuildJobResponse", string(data)}, " ")
}

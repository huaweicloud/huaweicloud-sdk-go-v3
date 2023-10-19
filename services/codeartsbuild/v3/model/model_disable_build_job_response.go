package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableBuildJobResponse Response Object
type DisableBuildJobResponse struct {

	// 返回结果状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableBuildJobResponse struct{}"
	}

	return strings.Join([]string{"DisableBuildJobResponse", string(data)}, " ")
}

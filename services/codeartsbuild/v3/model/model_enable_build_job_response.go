package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableBuildJobResponse Response Object
type EnableBuildJobResponse struct {

	// 返回结果状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnableBuildJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableBuildJobResponse struct{}"
	}

	return strings.Join([]string{"EnableBuildJobResponse", string(data)}, " ")
}

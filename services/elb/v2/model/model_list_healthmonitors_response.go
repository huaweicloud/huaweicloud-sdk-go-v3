package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHealthmonitorsResponse struct {
	// 健康检查对象的列表

	Healthmonitors *[]HealthmonitorResp `json:"healthmonitors,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListHealthmonitorsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHealthmonitorsResponse struct{}"
	}

	return strings.Join([]string{"ListHealthmonitorsResponse", string(data)}, " ")
}

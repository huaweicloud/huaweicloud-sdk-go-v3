package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServicesResponse Response Object
type ListServicesResponse struct {

	// 满足条件的端点个数
	Count *int64 `json:"count,omitempty"`

	// 服务列表
	Services       *[]ServiceRespDetail `json:"services,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListServicesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServicesResponse struct{}"
	}

	return strings.Join([]string{"ListServicesResponse", string(data)}, " ")
}

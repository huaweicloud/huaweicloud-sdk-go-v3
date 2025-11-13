package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUsageDataResponse Response Object
type ShowUsageDataResponse struct {

	// 资源总数。
	Count *int32 `json:"count,omitempty"`

	// 资源用量列表
	TimeList *[]TimeResourceUsageInfo `json:"time_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUsageDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUsageDataResponse struct{}"
	}

	return strings.Join([]string{"ShowUsageDataResponse", string(data)}, " ")
}

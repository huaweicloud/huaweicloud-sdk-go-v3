package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatusAppResponse Response Object
type ShowStatusAppResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 应用列表
	AppList *[]ShowStatusAppItem `json:"app_list,omitempty"`

	// 总访问次数
	TotalAccessCount *int64 `json:"total_access_count,omitempty"`

	// 成功访问次数
	SuccessAccessCount *int64 `json:"success_access_count,omitempty"`
	HttpStatusCode     int    `json:"-"`
}

func (o ShowStatusAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusAppResponse struct{}"
	}

	return strings.Join([]string{"ShowStatusAppResponse", string(data)}, " ")
}

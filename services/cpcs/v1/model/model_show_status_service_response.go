package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStatusServiceResponse Response Object
type ShowStatusServiceResponse struct {

	// 资源名称
	MetricName *string `json:"metric_name,omitempty"`

	// 服务列表
	ServerList     *[]ShowStatusServiceResponseBodyServerList `json:"server_list,omitempty"`
	HttpStatusCode int                                        `json:"-"`
}

func (o ShowStatusServiceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStatusServiceResponse struct{}"
	}

	return strings.Join([]string{"ShowStatusServiceResponse", string(data)}, " ")
}

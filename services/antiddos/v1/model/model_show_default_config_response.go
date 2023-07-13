package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultConfigResponse Response Object
type ShowDefaultConfigResponse struct {

	// 是否开启L7层防护
	EnableL7 *bool `json:"enable_L7,omitempty"`

	// 流量分段ID，取值范围：1～9
	TrafficPosId *int64 `json:"traffic_pos_id,omitempty"`

	// HTTP请求数分段ID，取值范围：1～15
	HttpRequestPosId *int64 `json:"http_request_pos_id,omitempty"`

	// 清洗时访问限制分段ID，取值范围：1～8
	CleaningAccessPosId *int64 `json:"cleaning_access_pos_id,omitempty"`

	// 应用类型ID，可选取值： - 0 - 1
	AppTypeId      *int64 `json:"app_type_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDefaultConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultConfigResponse", string(data)}, " ")
}

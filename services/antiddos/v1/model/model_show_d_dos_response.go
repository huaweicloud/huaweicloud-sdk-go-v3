package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowDDosResponse struct {
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

func (o ShowDDosResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowDDosResponse struct{}"
	}

	return strings.Join([]string{"ShowDDosResponse", string(data)}, " ")
}

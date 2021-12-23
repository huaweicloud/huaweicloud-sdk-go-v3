package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAntiDDosServiceRequestBody struct {
	// 应用类型ID，可选取值： - 0 - 1

	AppTypeId int32 `json:"app_type_id"`
	// 清洗时访问限制分段ID，取值范围：1～8

	CleaningAccessPosId int32 `json:"cleaning_access_pos_id"`
	// 是否开启L7层防护

	EnableL7 bool `json:"enable_L7"`
	// HTTP请求数分段ID，取值范围：1～15

	HttpRequestPosId int32 `json:"http_request_pos_id"`
	// 流量分段ID，取值范围：1～9

	TrafficPosId int32 `json:"traffic_pos_id"`
}

func (o UpdateAntiDDosServiceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAntiDDosServiceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAntiDDosServiceRequestBody", string(data)}, " ")
}

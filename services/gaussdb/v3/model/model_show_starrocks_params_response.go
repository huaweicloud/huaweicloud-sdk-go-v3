package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStarrocksParamsResponse Response Object
type ShowStarrocksParamsResponse struct {
	Configurations *ParameterConfigurationInfo `json:"configurations,omitempty"`

	// 参数信息的总数。
	TotalCount *int64 `json:"total_count,omitempty"`

	// 参数对象。
	ParameterValues *[]ParameterValuesInfo `json:"parameter_values,omitempty"`
	HttpStatusCode  int                    `json:"-"`
}

func (o ShowStarrocksParamsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStarrocksParamsResponse struct{}"
	}

	return strings.Join([]string{"ShowStarrocksParamsResponse", string(data)}, " ")
}

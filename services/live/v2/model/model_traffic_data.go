package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrafficData struct {

	// 采样值，单位为byte。
	Value *int64 `json:"value,omitempty"`

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time *string `json:"time,omitempty"`
}

func (o TrafficData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficData struct{}"
	}

	return strings.Join([]string{"TrafficData", string(data)}, " ")
}

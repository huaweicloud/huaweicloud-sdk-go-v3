package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEventDataResponse struct {
	// 配置信息列表。如果不存在对应的配置信息，则datapoints为空数组[]。

	Datapoints     *[]EventDataInfo `json:"datapoints,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowEventDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEventDataResponse struct{}"
	}

	return strings.Join([]string{"ShowEventDataResponse", string(data)}, " ")
}

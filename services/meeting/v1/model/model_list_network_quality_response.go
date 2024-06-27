package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworkQualityResponse Response Object
type ListNetworkQualityResponse struct {

	// 结果码
	ReturnCode *int32 `json:"returnCode,omitempty"`

	// 结果描述
	ReturnDesc *string `json:"returnDesc,omitempty"`

	// 会场Qos列表
	QosList        *[]UserQos `json:"qosList,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListNetworkQualityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworkQualityResponse struct{}"
	}

	return strings.Join([]string{"ListNetworkQualityResponse", string(data)}, " ")
}

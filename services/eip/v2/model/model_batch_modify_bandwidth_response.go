package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchModifyBandwidthResponse Response Object
type BatchModifyBandwidthResponse struct {

	// 成功资源
	SuccessResources *[]SuccessResources `json:"success_resources,omitempty"`

	// 失败资源
	FailureResources *[]FailureResources `json:"failure_resources,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o BatchModifyBandwidthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyBandwidthResponse struct{}"
	}

	return strings.Join([]string{"BatchModifyBandwidthResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrDeletePublishRecordForApiV2Request Request Object
type CreateOrDeletePublishRecordForApiV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ApiActionInfo `json:"body,omitempty"`
}

func (o CreateOrDeletePublishRecordForApiV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrDeletePublishRecordForApiV2Request struct{}"
	}

	return strings.Join([]string{"CreateOrDeletePublishRecordForApiV2Request", string(data)}, " ")
}

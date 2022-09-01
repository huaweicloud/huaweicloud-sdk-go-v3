package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLiveDataQuotaV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`
}

func (o ListLiveDataQuotaV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLiveDataQuotaV2Request struct{}"
	}

	return strings.Join([]string{"ListLiveDataQuotaV2Request", string(data)}, " ")
}

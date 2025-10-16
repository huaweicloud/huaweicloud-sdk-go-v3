package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDistributionRequest Request Object
type ListDistributionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ListDistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDistributionRequest struct{}"
	}

	return strings.Join([]string{"ListDistributionRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceDistributionRequest Request Object
type ShowInstanceDistributionRequest struct {

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`
}

func (o ShowInstanceDistributionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceDistributionRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceDistributionRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDistributionRequestBody 分发服务器信息。
type CreateDistributionRequestBody struct {

	// 配置为分发服务器的实例id。
	DistributorInstanceId string `json:"distributor_instance_id"`
}

func (o CreateDistributionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDistributionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDistributionRequestBody", string(data)}, " ")
}

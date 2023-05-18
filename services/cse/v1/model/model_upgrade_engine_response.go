package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpgradeEngineResponse struct {

	// 创建的微服务引擎专享版ID
	Id *string `json:"id,omitempty"`

	// 创建的微服务引擎专享版名称
	Name *string `json:"name,omitempty"`

	// 微服务引擎专享版执行的任务ID
	JobId          *int32 `json:"jobId,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpgradeEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeEngineResponse struct{}"
	}

	return strings.Join([]string{"UpgradeEngineResponse", string(data)}, " ")
}

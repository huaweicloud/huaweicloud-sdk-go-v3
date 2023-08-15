package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEngineResponse Response Object
type CreateEngineResponse struct {

	// 创建的微服务引擎专享版ID
	Id *string `json:"id,omitempty"`

	// 创建的微服务引擎专享版名称
	Name *string `json:"name,omitempty"`

	// 微服务引擎专享版执行的任务ID
	JobId          *int32 `json:"jobId,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEngineResponse struct{}"
	}

	return strings.Join([]string{"CreateEngineResponse", string(data)}, " ")
}

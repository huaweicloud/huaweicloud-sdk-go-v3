package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateEngineResponse struct {

	// 创建的微服务引擎专享版ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 创建的微服务引擎专享版名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 微服务引擎专享版执行的任务ID
	JobId          *int32 `json:"job_id,omitempty" xml:"job_id"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateEngineResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEngineResponse struct{}"
	}

	return strings.Join([]string{"CreateEngineResponse", string(data)}, " ")
}

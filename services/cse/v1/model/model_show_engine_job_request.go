package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEngineJobRequest Request Object
type ShowEngineJobRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 微服务引擎ID。
	EngineId string `json:"engine_id"`

	// 任务ID。
	JobId string `json:"job_id"`
}

func (o ShowEngineJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineJobRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineJobRequest", string(data)}, " ")
}

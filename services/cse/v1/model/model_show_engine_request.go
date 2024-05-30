package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEngineRequest Request Object
type ShowEngineRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 微服务引擎ID。
	EngineId string `json:"engine_id"`
}

func (o ShowEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEngineRequest struct{}"
	}

	return strings.Join([]string{"ShowEngineRequest", string(data)}, " ")
}

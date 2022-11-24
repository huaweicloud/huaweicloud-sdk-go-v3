package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RetryEngineRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 引擎id
	EngineId string `json:"engine_id"`

	Body *EngineAdditionalActionReq `json:"body,omitempty"`
}

func (o RetryEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryEngineRequest struct{}"
	}

	return strings.Join([]string{"RetryEngineRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResizeEngineRequest Request Object
type ResizeEngineRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 该字段内容填为 \"application/json;charset=UTF-8\"
	ContentType string `json:"Content-Type"`

	// 该字段内容填为 \"application/json\"
	Accept string `json:"Accept"`

	// 引擎id
	EngineId string `json:"engine_id"`

	Body *EngineModifyReq `json:"body,omitempty"`
}

func (o ResizeEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeEngineRequest struct{}"
	}

	return strings.Join([]string{"ResizeEngineRequest", string(data)}, " ")
}

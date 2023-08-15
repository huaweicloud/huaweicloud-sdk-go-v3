package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEngineRequest Request Object
type CreateEngineRequest struct {

	// 如果不带则默认企业项目为\"default\"，ID为\"0\"
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	Body *EngineCreateReq `json:"body,omitempty"`
}

func (o CreateEngineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEngineRequest struct{}"
	}

	return strings.Join([]string{"CreateEngineRequest", string(data)}, " ")
}

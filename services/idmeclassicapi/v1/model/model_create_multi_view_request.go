package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMultiViewRequest Request Object
type CreateMultiViewRequest struct {

	// 应用ID。
	Identifier string `json:"identifier"`

	Body *RdmParamVoMultiViewModelDto `json:"body,omitempty"`
}

func (o CreateMultiViewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiViewRequest struct{}"
	}

	return strings.Join([]string{"CreateMultiViewRequest", string(data)}, " ")
}

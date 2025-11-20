package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceConfigurationsRequest Request Object
type CreateResourceConfigurationsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *CreateResourceConfigurationsReqBody `json:"body,omitempty"`
}

func (o CreateResourceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"CreateResourceConfigurationsRequest", string(data)}, " ")
}

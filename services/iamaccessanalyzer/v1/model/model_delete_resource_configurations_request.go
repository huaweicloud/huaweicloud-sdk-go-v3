package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceConfigurationsRequest Request Object
type DeleteResourceConfigurationsRequest struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	Body *DeleteResourceConfigurationReqBody `json:"body,omitempty"`
}

func (o DeleteResourceConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"DeleteResourceConfigurationsRequest", string(data)}, " ")
}

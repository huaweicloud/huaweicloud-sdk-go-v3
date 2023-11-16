package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOpenSourceStrategyResponse Response Object
type DeleteOpenSourceStrategyResponse struct {

	// 状态
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteOpenSourceStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOpenSourceStrategyResponse struct{}"
	}

	return strings.Join([]string{"DeleteOpenSourceStrategyResponse", string(data)}, " ")
}

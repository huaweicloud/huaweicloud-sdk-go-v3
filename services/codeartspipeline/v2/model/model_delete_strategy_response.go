package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStrategyResponse Response Object
type DeleteStrategyResponse struct {

	// 状态
	Status         *bool `json:"status,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o DeleteStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStrategyResponse struct{}"
	}

	return strings.Join([]string{"DeleteStrategyResponse", string(data)}, " ")
}

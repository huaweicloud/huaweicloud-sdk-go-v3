package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWarmPoolNewResponse Response Object
type ShowWarmPoolNewResponse struct {
	WarmPool       *WarmPoolInfo `json:"warm_pool,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowWarmPoolNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWarmPoolNewResponse struct{}"
	}

	return strings.Join([]string{"ShowWarmPoolNewResponse", string(data)}, " ")
}

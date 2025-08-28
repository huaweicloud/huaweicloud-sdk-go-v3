package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWarmPoolResponse Response Object
type ShowWarmPoolResponse struct {
	WarmPool       *WarmPoolInfo `json:"warm_pool,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowWarmPoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWarmPoolResponse struct{}"
	}

	return strings.Join([]string{"ShowWarmPoolResponse", string(data)}, " ")
}

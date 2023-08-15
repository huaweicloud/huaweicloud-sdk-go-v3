package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceStatusInfoResponse Response Object
type ShowInstanceStatusInfoResponse struct {
	Result *InstanceStatusResponse `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceStatusInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusInfoResponse", string(data)}, " ")
}

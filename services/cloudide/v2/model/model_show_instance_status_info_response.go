package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceStatusInfoResponse struct {
	Result *InstanceStatusResponse `json:"result,omitempty" xml:"result"`

	// 状态
	Status         *string `json:"status,omitempty" xml:"status"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceStatusInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceStatusInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceStatusInfoResponse", string(data)}, " ")
}

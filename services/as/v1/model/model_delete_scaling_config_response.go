package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteScalingConfigResponse Response Object
type DeleteScalingConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingConfigResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingConfigResponse", string(data)}, " ")
}

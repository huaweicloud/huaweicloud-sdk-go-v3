package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteScalingInstanceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteScalingInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteScalingInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteScalingInstanceResponse", string(data)}, " ")
}

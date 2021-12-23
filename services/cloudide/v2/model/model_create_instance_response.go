package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateInstanceResponse struct {
	Result *InstancesResponseInstancesVoResult `json:"result,omitempty"`
	// 状态

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"CreateInstanceResponse", string(data)}, " ")
}

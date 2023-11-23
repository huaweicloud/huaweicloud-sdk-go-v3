package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMasterSlavePoolRequestBody This is a auto create Body Object
type CreateMasterSlavePoolRequestBody struct {
	Pool *CreateMasterSlavePoolOption `json:"pool"`
}

func (o CreateMasterSlavePoolRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMasterSlavePoolRequestBody struct{}"
	}

	return strings.Join([]string{"CreateMasterSlavePoolRequestBody", string(data)}, " ")
}

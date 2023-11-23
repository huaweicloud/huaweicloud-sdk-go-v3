package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMasterSlavePoolResponse Response Object
type DeleteMasterSlavePoolResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteMasterSlavePoolResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMasterSlavePoolResponse struct{}"
	}

	return strings.Join([]string{"DeleteMasterSlavePoolResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSubnetRequest struct {

	// 子网ID。
	SubnetId string `json:"subnet_id" xml:"subnet_id"`
}

func (o DeleteSubnetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubnetRequest struct{}"
	}

	return strings.Join([]string{"DeleteSubnetRequest", string(data)}, " ")
}

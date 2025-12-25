package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubnetId 子网ID
type SubnetId struct {
}

func (o SubnetId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetId struct{}"
	}

	return strings.Join([]string{"SubnetId", string(data)}, " ")
}

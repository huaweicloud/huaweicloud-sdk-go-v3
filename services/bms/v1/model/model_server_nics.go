package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 网卡Port ID
type ServerNics struct {
	Id *string `json:"id,omitempty"`
}

func (o ServerNics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNics struct{}"
	}

	return strings.Join([]string{"ServerNics", string(data)}, " ")
}

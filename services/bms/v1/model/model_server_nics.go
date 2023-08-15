package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerNics 网卡Port ID
type ServerNics struct {
	Id string `json:"id"`
}

func (o ServerNics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNics struct{}"
	}

	return strings.Join([]string{"ServerNics", string(data)}, " ")
}

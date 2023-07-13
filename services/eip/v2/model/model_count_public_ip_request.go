package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountPublicIpRequest Request Object
type CountPublicIpRequest struct {
}

func (o CountPublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPublicIpRequest struct{}"
	}

	return strings.Join([]string{"CountPublicIpRequest", string(data)}, " ")
}

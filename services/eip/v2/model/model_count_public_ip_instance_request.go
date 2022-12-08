package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CountPublicIpInstanceRequest struct {
}

func (o CountPublicIpInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountPublicIpInstanceRequest struct{}"
	}

	return strings.Join([]string{"CountPublicIpInstanceRequest", string(data)}, " ")
}

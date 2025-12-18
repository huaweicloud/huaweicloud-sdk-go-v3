package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetLongAkskConfigRequest Request Object
type GetLongAkskConfigRequest struct {
}

func (o GetLongAkskConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetLongAkskConfigRequest struct{}"
	}

	return strings.Join([]string{"GetLongAkskConfigRequest", string(data)}, " ")
}

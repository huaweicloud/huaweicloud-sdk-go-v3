package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InitializePlatformRequest Request Object
type InitializePlatformRequest struct {
}

func (o InitializePlatformRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InitializePlatformRequest struct{}"
	}

	return strings.Join([]string{"InitializePlatformRequest", string(data)}, " ")
}

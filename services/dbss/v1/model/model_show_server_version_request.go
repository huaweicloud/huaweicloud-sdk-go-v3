package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerVersionRequest Request Object
type ShowServerVersionRequest struct {
}

func (o ShowServerVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowServerVersionRequest", string(data)}, " ")
}

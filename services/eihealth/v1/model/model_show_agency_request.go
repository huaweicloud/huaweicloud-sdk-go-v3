package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAgencyRequest Request Object
type ShowAgencyRequest struct {
}

func (o ShowAgencyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAgencyRequest struct{}"
	}

	return strings.Join([]string{"ShowAgencyRequest", string(data)}, " ")
}

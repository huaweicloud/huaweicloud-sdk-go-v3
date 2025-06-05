package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultBuildParametersRequest Request Object
type ShowDefaultBuildParametersRequest struct {
}

func (o ShowDefaultBuildParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultBuildParametersRequest struct{}"
	}

	return strings.Join([]string{"ShowDefaultBuildParametersRequest", string(data)}, " ")
}

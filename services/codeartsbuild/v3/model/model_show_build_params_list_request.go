package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildParamsListRequest Request Object
type ShowBuildParamsListRequest struct {
}

func (o ShowBuildParamsListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildParamsListRequest struct{}"
	}

	return strings.Join([]string{"ShowBuildParamsListRequest", string(data)}, " ")
}

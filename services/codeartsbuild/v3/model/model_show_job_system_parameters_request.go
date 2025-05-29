package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobSystemParametersRequest Request Object
type ShowJobSystemParametersRequest struct {
}

func (o ShowJobSystemParametersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobSystemParametersRequest struct{}"
	}

	return strings.Join([]string{"ShowJobSystemParametersRequest", string(data)}, " ")
}

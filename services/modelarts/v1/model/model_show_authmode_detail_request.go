package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuthmodeDetailRequest Request Object
type ShowAuthmodeDetailRequest struct {
}

func (o ShowAuthmodeDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuthmodeDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAuthmodeDetailRequest", string(data)}, " ")
}

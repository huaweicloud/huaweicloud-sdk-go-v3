package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteGlobalEipRequest Request Object
type DeleteGlobalEipRequest struct {
	GlobalEipId string `json:"global_eip_id"`
}

func (o DeleteGlobalEipRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGlobalEipRequest struct{}"
	}

	return strings.Join([]string{"DeleteGlobalEipRequest", string(data)}, " ")
}

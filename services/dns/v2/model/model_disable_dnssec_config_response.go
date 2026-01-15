package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDnssecConfigResponse Response Object
type DisableDnssecConfigResponse struct {

	// **参数解释：** 状态。 **取值范围：** DISABLE：关闭
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableDnssecConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDnssecConfigResponse struct{}"
	}

	return strings.Join([]string{"DisableDnssecConfigResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAddressSetsResponse Response Object
type BatchDeleteAddressSetsResponse struct {

	// **参数解释**： 地址组信息 **取值范围**： 不涉及
	Data           *[]BatchDeleteAddressSetsRespData `json:"data,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o BatchDeleteAddressSetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAddressSetsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAddressSetsResponse", string(data)}, " ")
}

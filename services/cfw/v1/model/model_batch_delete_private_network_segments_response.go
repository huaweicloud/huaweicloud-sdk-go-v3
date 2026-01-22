package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePrivateNetworkSegmentsResponse Response Object
type BatchDeletePrivateNetworkSegmentsResponse struct {

	// **参数解释**： 防火墙ID **取值范围**： 不涉及
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchDeletePrivateNetworkSegmentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePrivateNetworkSegmentsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeletePrivateNetworkSegmentsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetAvaliableZoneResponse Response Object
type GetAvaliableZoneResponse struct {

	// **参数解释**: 可用区列表信息
	Body           *[]GetAvailableZoneResponseBody `json:"body,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o GetAvaliableZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAvaliableZoneResponse struct{}"
	}

	return strings.Join([]string{"GetAvaliableZoneResponse", string(data)}, " ")
}

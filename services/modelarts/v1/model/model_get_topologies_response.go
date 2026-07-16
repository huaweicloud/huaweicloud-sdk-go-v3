package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetTopologiesResponse Response Object
type GetTopologiesResponse struct {

	// **参数解释**：server列表。
	Servers        *[]ServerPhyInfo `json:"servers,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o GetTopologiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTopologiesResponse struct{}"
	}

	return strings.Join([]string{"GetTopologiesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicEntitiesResponse Response Object
type ListLogicEntitiesResponse struct {

	// 资产列表。
	Body           *[]Entity `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListLogicEntitiesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicEntitiesResponse struct{}"
	}

	return strings.Join([]string{"ListLogicEntitiesResponse", string(data)}, " ")
}

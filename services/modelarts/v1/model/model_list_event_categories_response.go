package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEventCategoriesResponse Response Object
type ListEventCategoriesResponse struct {

	// **参数解释**：事件类型。
	EventCategories *[]EventCategoriesResp `json:"event-categories,omitempty"`
	HttpStatusCode  int                    `json:"-"`
}

func (o ListEventCategoriesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEventCategoriesResponse struct{}"
	}

	return strings.Join([]string{"ListEventCategoriesResponse", string(data)}, " ")
}

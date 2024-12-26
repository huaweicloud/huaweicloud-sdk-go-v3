package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowListOfEventSampleResponse Response Object
type ShowListOfEventSampleResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 本页数量
	Size *int32 `json:"size,omitempty"`

	// 对象列表
	Items          *[]EventSampleItemInfo `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowListOfEventSampleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowListOfEventSampleResponse struct{}"
	}

	return strings.Join([]string{"ShowListOfEventSampleResponse", string(data)}, " ")
}

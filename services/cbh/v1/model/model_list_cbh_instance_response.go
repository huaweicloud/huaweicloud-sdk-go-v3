package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCbhInstanceResponse Response Object
type ListCbhInstanceResponse struct {

	// 云堡垒机实例总数。
	Total *int32 `json:"total,omitempty"`

	QuotaDetail *QuotaDetail `json:"quotaDetail,omitempty"`

	// 云堡垒机实例列表信息。
	Instance       *[]InstanceDetail `json:"instance,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListCbhInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCbhInstanceResponse struct{}"
	}

	return strings.Join([]string{"ListCbhInstanceResponse", string(data)}, " ")
}

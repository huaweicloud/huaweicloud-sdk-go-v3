package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApplyResponse Response Object
type ListApplyResponse struct {

	// 符合条件的申请总数
	Total *int32 `json:"total,omitempty"`

	// 本次返回的申请列表
	Records        *[]RecordForApplyDetail `json:"records,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListApplyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApplyResponse struct{}"
	}

	return strings.Join([]string{"ListApplyResponse", string(data)}, " ")
}

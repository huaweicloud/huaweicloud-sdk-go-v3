package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBlockIpsResponse Response Object
type ListBlockIpsResponse struct {

	// 封堵列表响应体
	BlockingList *[]BlockListBlockingList `json:"blocking_list,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBlockIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockIpsResponse struct{}"
	}

	return strings.Join([]string{"ListBlockIpsResponse", string(data)}, " ")
}

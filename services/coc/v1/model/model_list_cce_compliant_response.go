package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceCompliantResponse Response Object
type ListCceCompliantResponse struct {

	// 总数
	Count *int64 `json:"count,omitempty"`

	// CCE信息列表
	CceInfoList    *[]CceInfo `json:"cce_info_list,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListCceCompliantResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceCompliantResponse struct{}"
	}

	return strings.Join([]string{"ListCceCompliantResponse", string(data)}, " ")
}

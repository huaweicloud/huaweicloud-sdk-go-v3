package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSystemUserWhiteListResponse Response Object
type ListSystemUserWhiteListResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 可继续添加的白名单数量
	RemainNum *int32 `json:"remain_num,omitempty"`

	// 白名单数量上限
	LimitNum *int32 `json:"limit_num,omitempty"`

	// 系统用户白名单详情
	DataList       *[]SystemUserWhiteListResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                                `json:"-"`
}

func (o ListSystemUserWhiteListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemUserWhiteListResponse struct{}"
	}

	return strings.Join([]string{"ListSystemUserWhiteListResponse", string(data)}, " ")
}

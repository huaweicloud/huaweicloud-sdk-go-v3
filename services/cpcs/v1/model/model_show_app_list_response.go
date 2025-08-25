package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAppListResponse Response Object
type ShowAppListResponse struct {

	// 满足条件的应用总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 应用列表
	Result         *[]AppInfo `json:"result,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowAppListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAppListResponse struct{}"
	}

	return strings.Join([]string{"ShowAppListResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSmallVersionResponse Response Object
type ListSmallVersionResponse struct {

	// 参数解释： 数据库版本信息列表。
	DataStores *[]DatabaseSmallVersion `json:"data_stores,omitempty"`

	// 参数解释： 查询总个数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListSmallVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSmallVersionResponse struct{}"
	}

	return strings.Join([]string{"ListSmallVersionResponse", string(data)}, " ")
}

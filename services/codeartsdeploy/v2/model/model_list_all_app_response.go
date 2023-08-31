package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllAppResponse Response Object
type ListAllAppResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 应用列表数组
	Result         *[]AppExecutionInfo `json:"result,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListAllAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllAppResponse struct{}"
	}

	return strings.Join([]string{"ListAllAppResponse", string(data)}, " ")
}

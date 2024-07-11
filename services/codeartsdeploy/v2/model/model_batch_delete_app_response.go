package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAppResponse Response Object
type BatchDeleteAppResponse struct {

	// 批量删除应用结果
	Result *[]AppDeleteResult `json:"result,omitempty"`

	// 批量删除应用总数
	TotalNum       *int32 `json:"total_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAppResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAppResponse", string(data)}, " ")
}

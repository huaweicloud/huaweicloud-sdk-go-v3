package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetObjectsResponse Response Object
type BatchSetObjectsResponse struct {

	// 总数
	AllCounts *int64 `json:"all_counts,omitempty"`

	// 批量对象选择响应列表
	Results        *[]DatabaseObjectResp `json:"results,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o BatchSetObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetObjectsResponse struct{}"
	}

	return strings.Join([]string{"BatchSetObjectsResponse", string(data)}, " ")
}

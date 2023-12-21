package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesByTagResponse Response Object
type ListInstancesByTagResponse struct {

	// 根据查询模式获取到的资源实例详情。
	Resources *[]Resources `json:"resources,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstancesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesByTagResponse struct{}"
	}

	return strings.Join([]string{"ListInstancesByTagResponse", string(data)}, " ")
}

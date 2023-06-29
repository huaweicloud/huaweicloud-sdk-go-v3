package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteResourcesResponse Response Object
type BatchDeleteResourcesResponse struct {

	// 成功删除的资源数目
	SucceedCount   *int32 `json:"succeed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o BatchDeleteResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteResourcesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteResourcesResponse", string(data)}, " ")
}

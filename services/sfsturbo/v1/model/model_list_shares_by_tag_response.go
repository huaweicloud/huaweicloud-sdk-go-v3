package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSharesByTagResponse Response Object
type ListSharesByTagResponse struct {

	// 通过标签查询文件系统的资源列表
	Resources *[]ListSharesByTagResource `json:"resources,omitempty"`

	// 返回的文件系统的总量
	TotalCount *int32 `json:"total_count,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSharesByTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSharesByTagResponse struct{}"
	}

	return strings.Join([]string{"ListSharesByTagResponse", string(data)}, " ")
}

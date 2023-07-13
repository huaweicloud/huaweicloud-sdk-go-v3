package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertResourceInstancesResponse Response Object
type ListCertResourceInstancesResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源实例列表，详情请参见TagResource字段数据结构说明。
	Resources      *[]TagResource `json:"resources,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListCertResourceInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertResourceInstancesResponse struct{}"
	}

	return strings.Join([]string{"ListCertResourceInstancesResponse", string(data)}, " ")
}

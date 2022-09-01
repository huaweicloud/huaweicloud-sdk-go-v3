package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDedicatedResourcesResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 专属资源信息列表。
	Resources      *[]ListDedicatedResourceResult `json:"resources,omitempty" xml:"resources"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListDedicatedResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDedicatedResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDedicatedResourcesResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClustersByTagsResponse struct {

	// 资源列表
	Resources *[]MrsResource `json:"resources,omitempty" xml:"resources"`

	// 资源总数
	TotalCount     *int32 `json:"total_count,omitempty" xml:"total_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListClustersByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListClustersByTagsResponse", string(data)}, " ")
}

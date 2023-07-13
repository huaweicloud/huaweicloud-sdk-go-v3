package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceResourcesResponse Response Object
type ListServiceResourcesResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源基本信息列表，具体请参见表3。
	Infos          *[]ServiceResourceInfo `json:"infos,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListServiceResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListServiceResourcesResponse", string(data)}, " ")
}

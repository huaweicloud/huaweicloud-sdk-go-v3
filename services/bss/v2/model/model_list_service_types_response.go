package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServiceTypesResponse Response Object
type ListServiceTypesResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 云服务类型信息列表，具体请参见表3。
	ServiceTypes   *[]ServiceTypes `json:"service_types,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListServiceTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceTypesResponse struct{}"
	}

	return strings.Join([]string{"ListServiceTypesResponse", string(data)}, " ")
}

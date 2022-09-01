package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListRegionsResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 区域列表
	DataCenterList *[]DataCenterV2Do `json:"data_center_list,omitempty" xml:"data_center_list"`
	HttpStatusCode int               `json:"-"`
}

func (o ListRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListRegionsResponse", string(data)}, " ")
}

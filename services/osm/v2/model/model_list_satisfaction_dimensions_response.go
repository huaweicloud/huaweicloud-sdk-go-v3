package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSatisfactionDimensionsResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 满意度分类列表
	SatisfactionDimensionList *[]SatisfactionDimensionSimpleInfoV2 `json:"satisfaction_dimension_list,omitempty" xml:"satisfaction_dimension_list"`
	HttpStatusCode            int                                  `json:"-"`
}

func (o ListSatisfactionDimensionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSatisfactionDimensionsResponse struct{}"
	}

	return strings.Join([]string{"ListSatisfactionDimensionsResponse", string(data)}, " ")
}

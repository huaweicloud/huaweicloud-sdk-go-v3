package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProblemTypesResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 问题类型列表
	IncidentBusinessTypeList *[]SimpleIncidentBusinessTypeV2 `json:"incident_business_type_list,omitempty" xml:"incident_business_type_list"`
	HttpStatusCode           int                             `json:"-"`
}

func (o ListProblemTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProblemTypesResponse struct{}"
	}

	return strings.Join([]string{"ListProblemTypesResponse", string(data)}, " ")
}

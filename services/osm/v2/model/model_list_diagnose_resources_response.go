package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDiagnoseResourcesResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 诊断记录列表
	Result         *[]DiagnoseResourceVo `json:"result,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListDiagnoseResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnoseResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListDiagnoseResourcesResponse", string(data)}, " ")
}

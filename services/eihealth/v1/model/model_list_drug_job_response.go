package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDrugJobResponse struct {

	// 作业列表
	Jobs *[]DrugJobDto `json:"jobs,omitempty"`

	// 作业总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDrugJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDrugJobResponse struct{}"
	}

	return strings.Join([]string{"ListDrugJobResponse", string(data)}, " ")
}

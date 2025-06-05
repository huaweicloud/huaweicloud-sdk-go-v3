package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserDrugJobResponse Response Object
type ListUserDrugJobResponse struct {

	// 作业列表
	Jobs *[]DrugJobDto `json:"jobs,omitempty"`

	// 作业总数
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUserDrugJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserDrugJobResponse struct{}"
	}

	return strings.Join([]string{"ListUserDrugJobResponse", string(data)}, " ")
}

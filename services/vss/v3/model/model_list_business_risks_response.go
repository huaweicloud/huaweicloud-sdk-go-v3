package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBusinessRisksResponse struct {
	// 业务风险总数

	Total *int32 `json:"total,omitempty"`
	// 业务风险列表

	Data           *[]BusinessRiskItem `json:"data,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListBusinessRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessRisksResponse struct{}"
	}

	return strings.Join([]string{"ListBusinessRisksResponse", string(data)}, " ")
}

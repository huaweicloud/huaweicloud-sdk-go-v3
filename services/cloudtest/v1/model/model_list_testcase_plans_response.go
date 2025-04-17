package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTestcasePlansResponse Response Object
type ListTestcasePlansResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueListTestcasePlanVo `json:"result,omitempty"`

	Error          *ApiError `json:"error,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTestcasePlansResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTestcasePlansResponse struct{}"
	}

	return strings.Join([]string{"ListTestcasePlansResponse", string(data)}, " ")
}

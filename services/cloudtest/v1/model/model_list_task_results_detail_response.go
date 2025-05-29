package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResultsDetailResponse Response Object
type ListTaskResultsDetailResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueListTaskResulDetailtVo `json:"result,omitempty"`

	Error          *ApiError `json:"error,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTaskResultsDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResultsDetailResponse struct{}"
	}

	return strings.Join([]string{"ListTaskResultsDetailResponse", string(data)}, " ")
}

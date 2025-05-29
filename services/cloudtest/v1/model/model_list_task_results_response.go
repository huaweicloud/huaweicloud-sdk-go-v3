package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskResultsResponse Response Object
type ListTaskResultsResponse struct {

	// success|error
	Status *string `json:"status,omitempty"`

	Result *ResultValueListTaskResultVo `json:"result,omitempty"`

	Error          *ApiError `json:"error,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListTaskResultsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskResultsResponse struct{}"
	}

	return strings.Join([]string{"ListTaskResultsResponse", string(data)}, " ")
}

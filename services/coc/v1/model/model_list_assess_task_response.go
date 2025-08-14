package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssessTaskResponse Response Object
type ListAssessTaskResponse struct {
	Data           *BaseQueryAssessTaskListResponseData `json:"data,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o ListAssessTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssessTaskResponse struct{}"
	}

	return strings.Join([]string{"ListAssessTaskResponse", string(data)}, " ")
}

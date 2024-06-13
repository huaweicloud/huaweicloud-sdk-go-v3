package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectTestCaseFieldsResponse Response Object
type ListProjectTestCaseFieldsResponse struct {
	Value          *[]ProjectTestCaseFieldVo `json:"value,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListProjectTestCaseFieldsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectTestCaseFieldsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectTestCaseFieldsResponse", string(data)}, " ")
}

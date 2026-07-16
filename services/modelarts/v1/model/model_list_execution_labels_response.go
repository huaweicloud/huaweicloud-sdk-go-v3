package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExecutionLabelsResponse Response Object
type ListExecutionLabelsResponse struct {

	// **参数解释：** 返回标签数组。
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListExecutionLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExecutionLabelsResponse struct{}"
	}

	return strings.Join([]string{"ListExecutionLabelsResponse", string(data)}, " ")
}

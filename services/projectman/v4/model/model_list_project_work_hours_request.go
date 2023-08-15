package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectWorkHoursRequest Request Object
type ListProjectWorkHoursRequest struct {
	Body *ListProjectWorkHoursRequestBody `json:"body,omitempty"`
}

func (o ListProjectWorkHoursRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectWorkHoursRequest struct{}"
	}

	return strings.Join([]string{"ListProjectWorkHoursRequest", string(data)}, " ")
}

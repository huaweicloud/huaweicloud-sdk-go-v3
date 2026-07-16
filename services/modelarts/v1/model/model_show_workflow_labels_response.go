package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowLabelsResponse Response Object
type ShowWorkflowLabelsResponse struct {

	// **参数解释：** 返回标签数组。
	Body           *[]string `json:"body,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowWorkflowLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowLabelsResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowLabelsResponse", string(data)}, " ")
}

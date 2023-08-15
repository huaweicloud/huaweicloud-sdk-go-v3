package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishWorkflowResponse Response Object
type PublishWorkflowResponse struct {

	// 资产id
	Id *string `json:"id,omitempty"`

	// 资产版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishWorkflowResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishWorkflowResponse struct{}"
	}

	return strings.Join([]string{"PublishWorkflowResponse", string(data)}, " ")
}

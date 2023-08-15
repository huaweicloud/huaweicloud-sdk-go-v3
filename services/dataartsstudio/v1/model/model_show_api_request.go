package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowApiRequest Request Object
type ShowApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// API ID
	ApiId string `json:"api_id"`
}

func (o ShowApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiRequest struct{}"
	}

	return strings.Join([]string{"ShowApiRequest", string(data)}, " ")
}

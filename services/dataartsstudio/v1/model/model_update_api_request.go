package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApiRequest Request Object
type UpdateApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// API ID
	ApiId string `json:"api_id"`

	Body *Api `json:"body,omitempty"`
}

func (o UpdateApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApiRequest struct{}"
	}

	return strings.Join([]string{"UpdateApiRequest", string(data)}, " ")
}

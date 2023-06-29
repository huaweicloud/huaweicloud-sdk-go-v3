package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishApiRequest Request Object
type PublishApiRequest struct {

	// 工作空间id
	Workspace *string `json:"workspace,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	Body *OpenApiParaForPublish `json:"body,omitempty"`
}

func (o PublishApiRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishApiRequest struct{}"
	}

	return strings.Join([]string{"PublishApiRequest", string(data)}, " ")
}

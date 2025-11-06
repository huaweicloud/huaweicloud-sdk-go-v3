package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebhookRepositoryDto struct {

	// **参数解释：** 仓库ID
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 仓库路径
	Namespace *string `json:"namespace,omitempty"`
}

func (o WebhookRepositoryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebhookRepositoryDto struct{}"
	}

	return strings.Join([]string{"WebhookRepositoryDto", string(data)}, " ")
}

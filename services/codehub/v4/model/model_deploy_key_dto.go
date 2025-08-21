package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployKeyDto **参数解释：** 部署密钥对象。
type DeployKeyDto struct {

	// **参数解释：** 部署密钥ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 部署密钥标题。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 部署密钥指纹。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o DeployKeyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployKeyDto struct{}"
	}

	return strings.Join([]string{"DeployKeyDto", string(data)}, " ")
}

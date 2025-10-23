package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployKeyValueDto 部署密钥
type DeployKeyValueDto struct {

	// **参数解释：** 待检查密钥值。 **取值范围：** 字符串长度不少于1，不超过5000。
	Key *string `json:"key,omitempty"`
}

func (o DeployKeyValueDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployKeyValueDto struct{}"
	}

	return strings.Join([]string{"DeployKeyValueDto", string(data)}, " ")
}

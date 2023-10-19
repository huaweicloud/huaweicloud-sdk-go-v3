package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAuthorisation 更新授权的详细信息。
type UpdateAuthorisation struct {

	// 实例名字。
	Name *string `json:"name,omitempty"`

	// 实例描述。不支持 <>。
	Description *string `json:"description,omitempty"`
}

func (o UpdateAuthorisation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAuthorisation struct{}"
	}

	return strings.Join([]string{"UpdateAuthorisation", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepositoryInheritSettingResponse Response Object
type UpdateRepositoryInheritSettingResponse struct {

	// 仓库继承设置列表
	Body           *[]RepositoryInheritSettingDto `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o UpdateRepositoryInheritSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryInheritSettingResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryInheritSettingResponse", string(data)}, " ")
}

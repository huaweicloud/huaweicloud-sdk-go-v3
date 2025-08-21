package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRepositoryInheritSettingResponse Response Object
type ShowRepositoryInheritSettingResponse struct {

	// 仓库继承设置列表
	Body           *[]RepositoryInheritSettingDto `json:"body,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowRepositoryInheritSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepositoryInheritSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowRepositoryInheritSettingResponse", string(data)}, " ")
}

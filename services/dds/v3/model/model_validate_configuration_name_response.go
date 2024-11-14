package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateConfigurationNameResponse Response Object
type ValidateConfigurationNameResponse struct {

	// 参数组名称是否存在。 - true：参数组名称存在 - false：参数组名称不存在
	IsExisted      *bool `json:"is_existed,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateConfigurationNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConfigurationNameResponse struct{}"
	}

	return strings.Join([]string{"ValidateConfigurationNameResponse", string(data)}, " ")
}

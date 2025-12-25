package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateConfigurationApplicationResponse Response Object
type CreateConfigurationApplicationResponse struct {

	// 创建失败实例列表
	ApplyFailList *[]ComponentConfiguration `json:"apply_fail_list,omitempty"`

	// 创建成功实例列表
	ApplySuccessList *[]ComponentConfiguration `json:"apply_success_list,omitempty"`
	HttpStatusCode   int                       `json:"-"`
}

func (o CreateConfigurationApplicationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConfigurationApplicationResponse struct{}"
	}

	return strings.Join([]string{"CreateConfigurationApplicationResponse", string(data)}, " ")
}

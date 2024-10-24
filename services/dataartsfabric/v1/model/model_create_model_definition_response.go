package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelDefinitionResponse Response Object
type CreateModelDefinitionResponse struct {

	// 模型ID，32~36位的英文、数字、短横组合
	Id *string `json:"id,omitempty"`

	// 模型版本ID，32~36位的英文、数字、短横组合，系统自动生成无法修改，输入不生效。
	VersionId *string `json:"version_id,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateModelDefinitionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelDefinitionResponse struct{}"
	}

	return strings.Join([]string{"CreateModelDefinitionResponse", string(data)}, " ")
}

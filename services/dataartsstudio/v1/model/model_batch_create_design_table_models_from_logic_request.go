package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateDesignTableModelsFromLogicRequest Request Object
type BatchCreateDesignTableModelsFromLogicRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 所属关系建模的模型ID，ID字符串。
	ModelId string `json:"model_id"`

	Body *WorkspaceTransformVo `json:"body,omitempty"`
}

func (o BatchCreateDesignTableModelsFromLogicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDesignTableModelsFromLogicRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateDesignTableModelsFromLogicRequest", string(data)}, " ")
}

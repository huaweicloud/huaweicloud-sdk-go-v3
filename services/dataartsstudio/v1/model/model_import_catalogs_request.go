package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCatalogsRequest Request Object
type ImportCatalogsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 需要执行的动作。 枚举值：   - start-import: 开始导入
	ActionId string `json:"action-id"`

	// 是否需要覆盖更新已有的主题。
	SkipExist *bool `json:"skip-exist,omitempty"`

	Body *ImportCatalogsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportCatalogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCatalogsRequest struct{}"
	}

	return strings.Join([]string{"ImportCatalogsRequest", string(data)}, " ")
}

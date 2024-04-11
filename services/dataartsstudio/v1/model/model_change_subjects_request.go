package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeSubjectsRequest Request Object
type ChangeSubjectsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	Body *CatalogLevelVoList `json:"body,omitempty"`
}

func (o ChangeSubjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeSubjectsRequest struct{}"
	}

	return strings.Join([]string{"ChangeSubjectsRequest", string(data)}, " ")
}

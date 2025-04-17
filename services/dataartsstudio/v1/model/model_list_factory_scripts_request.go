package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactoryScriptsRequest Request Object
type ListFactoryScriptsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 分页返回结果，指定每页最大记录数，范围[1,100]。 默认值：10。
	Limit *int32 `json:"limit,omitempty"`

	// 分页的起始页，默认值为0。取值范围大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 脚本名称
	ScriptName *string `json:"script_name,omitempty"`
}

func (o ListFactoryScriptsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactoryScriptsRequest struct{}"
	}

	return strings.Join([]string{"ListFactoryScriptsRequest", string(data)}, " ")
}

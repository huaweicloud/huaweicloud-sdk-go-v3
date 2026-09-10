package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRealTimeJobDetailsRequest Request Object
type ShowRealTimeJobDetailsRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 作业名称。
	JobName string `json:"job_name"`
}

func (o ShowRealTimeJobDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRealTimeJobDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowRealTimeJobDetailsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudPhoneServerRequestBodyExtendParam 扩展字段。
type ChangeCloudPhoneServerRequestBodyExtendParam struct {

	// 企业项目ID。 该字段不传（或传为字符串“0”），则将资源绑定给默认企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ChangeCloudPhoneServerRequestBodyExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudPhoneServerRequestBodyExtendParam struct{}"
	}

	return strings.Join([]string{"ChangeCloudPhoneServerRequestBodyExtendParam", string(data)}, " ")
}

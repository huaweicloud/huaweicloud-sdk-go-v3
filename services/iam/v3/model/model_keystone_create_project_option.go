package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 项目信息。
type KeystoneCreateProjectOption struct {
	// 项目名称。必须以存在的\"区域ID_\"开头，长度小于等于64字符。例如区域“华北-北京一”的区域ID为“cn-north-1”，在其下创建项目时，项目名应填“cn-north-1_IAMProject”

	Name string `json:"name"`
	// 区域对应的项目ID，例如区域“华北-北京一”区域对应的项目ID为：04dd42abe48026ad2fa3c01ad7fa.....，获取方式请参见：[获取账号、IAM用户、项目、用户组、委托的名称和ID](https://support.huaweicloud.com/api-iam/iam_17_0002.html)。

	ParentId string `json:"parent_id"`
	// 项目所属账号ID。

	DomainId *string `json:"domain_id,omitempty"`
	// 项目描述信息，长度小于等于255字符。

	Description *string `json:"description,omitempty"`
}

func (o KeystoneCreateProjectOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneCreateProjectOption struct{}"
	}

	return strings.Join([]string{"KeystoneCreateProjectOption", string(data)}, " ")
}

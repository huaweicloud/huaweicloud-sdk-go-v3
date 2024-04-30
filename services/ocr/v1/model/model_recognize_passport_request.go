package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecognizePassportRequest Request Object
type RecognizePassportRequest struct {

	// 企业项目ID。OCR支持通过企业项目管理（EPS）对不同用户组和用户的资源使用，进行分账。 获取方法：[进入“[企业项目管理](https://console.huaweicloud.com/eps/?region=cn-north-4#/projects/list)”页面,](tag:hc)[进入“[企业项目管理](https://console-intl.huaweicloud.com/eps/?region=ap-southeast-1#/projects/list)”页面,](tag:hk)单击企业项目名称，在企业项目详情页获取Enterprise-Project-Id（企业项目ID）。 企业项目创建步骤请参见用户指南。 > 说明： 创建企业项目后，在传参时，有以下三类场景。 - 携带正确的ID，正常使用OCR服务，账单的企业项目会被分类到企业ID对应的企业项目中。 - 携带格式正确但不存在的ID，正常使用OCR服务，账单的企业项目会显示对应不存在的企业项目ID。 - 不携带ID或格式错误ID（包含特殊字符等），正常使用OCR服务，账单的企业项目会被分类到\"default\"中。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *PassportRequestBody `json:"body,omitempty"`
}

func (o RecognizePassportRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecognizePassportRequest struct{}"
	}

	return strings.Join([]string{"RecognizePassportRequest", string(data)}, " ")
}

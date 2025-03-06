package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNoticeRuleRequest Request Object
type CreateNoticeRuleRequest struct {

	// 企业项目ID。  - 创建环境时，环境会绑定企业项目ID。      - 最大长度36字节，带“-”连字符的UUID格式，或者是字符串“0”。     - 该字段不传（或传为字符串“0”）时，则查询默认企业项目下的资源。  > 关于企业项目ID的获取及企业项目特性的详细信息，请参见《[企业管理服务用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)》。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	Body *CreateNoticeRuleReq `json:"body,omitempty"`
}

func (o CreateNoticeRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNoticeRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateNoticeRuleRequest", string(data)}, " ")
}

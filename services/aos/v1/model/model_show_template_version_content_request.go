package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTemplateVersionContentRequest Request Object
type ShowTemplateVersionContentRequest struct {

	// 用户指定的，对于此请求的唯一ID，用于定位某个请求，推荐使用UUID
	ClientRequestId string `json:"Client-Request-Id"`

	// 用户希望创建的模板名称
	TemplateName string `json:"template_name"`

	// 模板版本ID，以大写V开头，每次创建模板版本，模板版本ID数字部分会自增加一
	VersionId string `json:"version_id"`

	// 模板的ID。当template_id存在时，模板服务会检查template_id是否和template_name匹配，不匹配会返回400
	TemplateId *string `json:"template_id,omitempty"`

	// 允许访问资源栈模板的source vpc id列表， source vpc id应仅包含小写字母、数字或中划线。
	AccessControlSourceVpcIds *[]string `json:"access_control_source_vpc_ids,omitempty"`

	// 允许访问资源栈模板的source ip列表，source ip应是具有CIDR表示法且带有子网掩码的IPv4地址。
	AccessControlSourceIps *[]string `json:"access_control_source_ips,omitempty"`
}

func (o ShowTemplateVersionContentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTemplateVersionContentRequest struct{}"
	}

	return strings.Join([]string{"ShowTemplateVersionContentRequest", string(data)}, " ")
}

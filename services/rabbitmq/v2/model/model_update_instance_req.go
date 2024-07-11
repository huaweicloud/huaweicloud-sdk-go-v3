package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceReq struct {

	// 实例名称。  由英文字符开头，只能由英文字母、数字、中划线组成，长度为4~64的字符。
	Name *string `json:"name,omitempty"`

	// 实例的描述信息。  长度不超过1024的字符串。[且字符串不能包含\">\"与\"<\"，字符串首字符不能为\"=\",\"+\",\"-\",\"@\"的全角和半角字符。](tag:hcs) > \\与\"在json报文中属于特殊字符，如果参数值中需要显示\\或者\"字符，请在字符前增加转义字符\\，比如\\\\或者\\\"。
	Description *string `json:"description,omitempty"`

	// 维护时间窗开始时间，格式为HH:mm:ss。   - 维护时间窗开始和结束时间必须为指定的时间段。   - 开始时间必须为22:00:00、02:00:00、06:00:00、10:00:00、14:00:00和18:00:00。   - 该参数不能单独为空，若该值为空，则结束时间也为空。系统分配一个默认开始时间02:00:00。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间，格式为HH:mm:ss。   - 维护时间窗开始和结束时间必须为指定的时间段。   - 结束时间在开始时间基础上加四个小时，即当开始时间为22:00:00时，结束时间为02:00:00。   - 该参数不能单独为空，若该值为空，则开始时间也为空。系统分配一个默认结束时间06:00:00。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 安全组ID。  获取方法如下：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// RabbitMQ实例是否开启公网访问功能。   - true：开启   - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// RabbitMQ实例绑定的弹性IP地址的id。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。  获取方法：登录弹性公网IP和带宽的控制台界面，在弹性公网IP的详情页面查的基本信息栏找ID。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 企业项目。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// ACL访问控制[（仅AMQP版本支持此参数）](tag:hws,hws_hk)。
	EnableAcl *bool `json:"enable_acl,omitempty"`
}

func (o UpdateInstanceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceReq struct{}"
	}

	return strings.Join([]string{"UpdateInstanceReq", string(data)}, " ")
}

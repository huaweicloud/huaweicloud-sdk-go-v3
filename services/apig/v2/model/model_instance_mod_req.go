package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceModReq struct {

	// 实例描述。支持除>和<以外的字符，长度为0~255。
	Description *string `json:"description,omitempty"`

	// 维护时间窗开始时间。时间格式为 xx:00:00，xx取值为02,06,10,14,18,22。  在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间。时间格式为 xx:00:00，与维护时间窗开始时间相差4个小时。  在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 实例名称。  中英文字符开头，只能由中英文字符、数字、中划线、下划线组成，长度为3~64。  > 中文字符必须为UTF-8或者unicode编码。
	InstanceName *string `json:"instance_name,omitempty"`

	// 指定实例所属的安全组。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询安全组列表”章节。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 终端节点服务的名称。  支持英文、数字、中划线、下划线，0~16个字符。  如果您填写该参数为空，系统生成的终端节点服务的名称为{region}.{service_id}。 如果您填写该参数，系统生成的终端节点服务的名称为{region}.{vpcep_service_name}.{service_id}。
	VpcepServiceName *string `json:"vpcep_service_name,omitempty"`
}

func (o InstanceModReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceModReq struct{}"
	}

	return strings.Join([]string{"InstanceModReq", string(data)}, " ")
}

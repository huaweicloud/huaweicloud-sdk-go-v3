package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceDetail 实例相关信息的集合
type CreateInstanceDetail struct {

	// DDM实例名称，命名要求如下。 - 长度为4-64个字符。 - 必须以字母开头。 - 可以包含字母、数字、中划线，不能包含其它特殊字符。
	Name string `json:"name"`

	// 规格ID。
	FlavorId string `json:"flavor_id"`

	// 节点个数。
	NodeNum int32 `json:"node_num"`

	// 引擎ID。
	EngineId string `json:"engine_id"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 可用区code。取值非空，请参见地区和终端节点(https://developer.huaweicloud.com/endpoint?DDM)。
	AvailableZones []string `json:"available_zones"`

	// 虚拟私有云的ID。
	VpcId string `json:"vpc_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id"`

	// 子网ID。
	SubnetId string `json:"subnet_id"`

	// 参数组ID.
	ParamGroupId *string `json:"param_group_id,omitempty"`

	// UTC时区。默认为UTC。取值范围：\"UTC\",\"UTC-12:00\",\"UTC-11:00\",\"UTC-10:00\",\"UTC-09:00\", \"UTC-08:00\", \"UTC-07:00\", \"UTC-06:00\", \"UTC-05:00\", \"UTC-04:00\", \"UTC-03:00\", \"UTC-02:00\", \"UTC-01:00\", \"UTC+01:00\", \"UTC+02:00\", \"UTC+03:00\", \"UTC+04:00\", \"UTC+05:00\", \"UTC+06:00\", \"UTC+07:00\", \"UTC+08:00\", \"UTC+09:00\", \"UTC+10:00\", \"UTC+11:00\", \"UTC+12:00\"
	TimeZone *string `json:"time_zone,omitempty"`

	// 管理员账号用户名。 - 长度为1-32个字符。 - 必须以字母开头。 - 可以包含字母，数字、下划线，不能包含其它特殊字符。
	AdminUserName *string `json:"admin_user_name,omitempty"`

	// 管理员账号密码。 - 长度为8~32位。 - 必须是大写字母（A~Z）、小写字母（a~z）、数字（0~9）、特殊字符~!@#%^*-_=+?的组合。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。
	AdminUserPassword *string `json:"admin_user_password,omitempty"`
}

func (o CreateInstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceDetail struct{}"
	}

	return strings.Join([]string{"CreateInstanceDetail", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 实例相关信息的集合
type CreateInstanceDetail struct {

	// DDM实例名称，命名要求如下。 - 长度为4-64个字符。 - 必须以字母开头。 - 可以包含字母、数字、中划线，不能包含其它特殊字符。
	Name string `json:"name" xml:"name"`

	// 规格ID。
	FlavorId string `json:"flavor_id" xml:"flavor_id"`

	// 节点个数。
	NodeNum int32 `json:"node_num" xml:"node_num"`

	// 引擎ID。
	EngineId string `json:"engine_id" xml:"engine_id"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 可用区code。取值非空，请参见地区和终端节点(https://developer.huaweicloud.com/endpoint?DDM)。
	AvailableZones []string `json:"available_zones" xml:"available_zones"`

	// 虚拟私有云的ID。
	VpcId string `json:"vpc_id" xml:"vpc_id"`

	// 安全组ID。
	SecurityGroupId string `json:"security_group_id" xml:"security_group_id"`

	// 子网ID。
	SubnetId string `json:"subnet_id" xml:"subnet_id"`

	// 参数组ID.
	ParamGroupId *string `json:"param_group_id,omitempty" xml:"param_group_id"`

	// UTC时区。默认为UTC。取值范围：\"UTC\",\"UTC-12:00\",\"UTC-11:00\",\"UTC-10:00\",\"UTC-09:00\", \"UTC-08:00\", \"UTC-07:00\", \"UTC-06:00\", \"UTC-05:00\", \"UTC-04:00\", \"UTC-03:00\", \"UTC-02:00\", \"UTC-01:00\", \"UTC+01:00\", \"UTC+02:00\", \"UTC+03:00\", \"UTC+04:00\", \"UTC+05:00\", \"UTC+06:00\", \"UTC+07:00\", \"UTC+08:00\", \"UTC+09:00\", \"UTC+10:00\", \"UTC+11:00\", \"UTC+12:00\"
	TimeZone *string `json:"time_zone,omitempty" xml:"time_zone"`
}

func (o CreateInstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceDetail struct{}"
	}

	return strings.Join([]string{"CreateInstanceDetail", string(data)}, " ")
}

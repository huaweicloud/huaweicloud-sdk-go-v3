package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SrCreateInstanceRspInstance 实例信息。
type SrCreateInstanceRspInstance struct {

	// StarRocks实例ID，严格匹配UUID规则。
	Id *string `json:"id,omitempty"`

	// 可用区码。
	AzCode *string `json:"az_code,omitempty"`

	// 可用区模式。  取值范围：  single：单可用区。  multi：多可用区。
	AzMode *string `json:"az_mode,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	Engine *SrCreateInstanceRspInstanceEngine `json:"engine,omitempty"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 子网ID。
	SubNetId *string `json:"sub_net_id,omitempty"`

	// 数据库用户。默认root。
	DbUser *string `json:"db_user,omitempty"`

	// 数据库端口号。默认3306。
	Port *int32 `json:"port,omitempty"`

	// 部署模式。
	HaMode *string `json:"ha_mode,omitempty"`

	PayInfo *SrCreateInstanceRspInstancePayInfo `json:"pay_info,omitempty"`

	// SSL开关。
	EnableSsl *bool `json:"enable_ssl,omitempty"`

	// 实例状态。
	Status *string `json:"status,omitempty"`

	// 实例所在区域。
	Region *string `json:"region,omitempty"`

	TagsInfo *SrCreateInstanceRspInstanceTagsInfo `json:"tags_info,omitempty"`
}

func (o SrCreateInstanceRspInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SrCreateInstanceRspInstance struct{}"
	}

	return strings.Join([]string{"SrCreateInstanceRspInstance", string(data)}, " ")
}

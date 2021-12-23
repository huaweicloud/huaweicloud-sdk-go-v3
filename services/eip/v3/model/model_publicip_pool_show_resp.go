package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 公网IP池详情
type PublicipPoolShowResp struct {
	// 公网IP池id

	Id *string `json:"id,omitempty"`
	// 公网IP池名字

	Name *string `json:"name,omitempty"`
	// 状态

	Status *string `json:"status,omitempty"`
	// 取值, spec_bgp(专属离散动态), spec_sbgp(专属离散静态)

	Type *PublicipPoolShowRespType `json:"type,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
	// 租户id

	ProjectId *string `json:"project_id,omitempty"`
	// 池子大小

	Size *int32 `json:"size,omitempty"`
	// 已经使用的ip数量

	Used *int32 `json:"used,omitempty"`
	// 公网IP池创建时间

	CreatedAt *string `json:"created_at,omitempty"`
	// 公网IP池更新时间

	UpdatedAt *string `json:"updated_at,omitempty"`

	BillingInfo *BillingInfoDict `json:"billing_info,omitempty"`
	// 功能说明：中心还是边缘。公网IP池取值为center

	PublicBorderGroup *string `json:"public_border_group,omitempty"`
	// 功能说明：是否共享

	Shared *bool `json:"shared,omitempty"`
	// 功能说明：是否公共池

	IsCommon *bool `json:"is_common,omitempty"`
	// 默认不显示。用户标签

	Tags *[]TagsInfo `json:"tags,omitempty"`
	// 功能说明：企业项目ID。最大长度36字节,带“-”连字符的UUID格式,或者是字符串“0”。创建弹性公网IP时,给弹性公网IP绑定企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 功能说明：表示此publicip可以加入的共享带宽类型列表，如果为空列表，则表示该           publicip不能加入任何共享带宽 约束：publicip只能加入到有该带宽类型的共享带宽中

	AllowShareBandwidthTypes *[]string `json:"allow_share_bandwidth_types,omitempty"`
}

func (o PublicipPoolShowResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublicipPoolShowResp struct{}"
	}

	return strings.Join([]string{"PublicipPoolShowResp", string(data)}, " ")
}

type PublicipPoolShowRespType struct {
	value string
}

type PublicipPoolShowRespTypeEnum struct {
	SPEC_BGP  PublicipPoolShowRespType
	SPEC_SBGP PublicipPoolShowRespType
}

func GetPublicipPoolShowRespTypeEnum() PublicipPoolShowRespTypeEnum {
	return PublicipPoolShowRespTypeEnum{
		SPEC_BGP: PublicipPoolShowRespType{
			value: "spec_bgp",
		},
		SPEC_SBGP: PublicipPoolShowRespType{
			value: "spec_sbgp",
		},
	}
}

func (c PublicipPoolShowRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublicipPoolShowRespType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

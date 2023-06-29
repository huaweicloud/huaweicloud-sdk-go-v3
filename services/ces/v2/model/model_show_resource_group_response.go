package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// ShowResourceGroupResponse Response Object
type ShowResourceGroupResponse struct {

	// 资源分组的名称
	GroupName *string `json:"group_name,omitempty"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId *string `json:"group_id,omitempty"`

	// 资源分组的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 资源分组归属企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源分组创建方式，取值只能为EPS（同步企业项目）,TAG（标签动态匹配）,Manual（手动添加）
	Type *ShowResourceGroupResponseType `json:"type,omitempty"`

	// 该资源分组内包含的资源来源的企业项目ID，type为EPS时必传
	AssociationEpIds *[]string `json:"association_ep_ids,omitempty"`

	// 标签动态匹配时的关联标签,type为TAG时必传
	Tags           *[]ResourceGroupTagRelation `json:"tags,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ShowResourceGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowResourceGroupResponse", string(data)}, " ")
}

type ShowResourceGroupResponseType struct {
	value string
}

type ShowResourceGroupResponseTypeEnum struct {
	EPS    ShowResourceGroupResponseType
	TAG    ShowResourceGroupResponseType
	MANUAL ShowResourceGroupResponseType
}

func GetShowResourceGroupResponseTypeEnum() ShowResourceGroupResponseTypeEnum {
	return ShowResourceGroupResponseTypeEnum{
		EPS: ShowResourceGroupResponseType{
			value: "EPS",
		},
		TAG: ShowResourceGroupResponseType{
			value: "TAG",
		},
		MANUAL: ShowResourceGroupResponseType{
			value: "Manual",
		},
	}
}

func (c ShowResourceGroupResponseType) Value() string {
	return c.value
}

func (c ShowResourceGroupResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResourceGroupResponseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

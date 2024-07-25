package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateWorkspaceResponseBodyWorkspaceAgencyList struct {

	// 委托空间所属项目id
	ProjectId *string `json:"project_id,omitempty"`

	// 空间委托id
	Id *string `json:"id,omitempty"`

	// 空间委托名称
	Name *string `json:"name,omitempty"`

	// 委托空间所属region id
	RegionId *string `json:"region_id,omitempty"`

	// THIS_ACCOUNT:本账号空间,CROSS_ACCOUNT:跨账号空间
	WorkspaceAttribution *CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution `json:"workspace_attribution,omitempty"`

	// 用户创建托管空间时使用的IAM委托版本，V3或者V5
	AgencyVersion *string `json:"agency_version,omitempty"`

	// 委托租户id
	DomainId *string `json:"domain_id,omitempty"`

	// 委托租户名称
	DomainName *string `json:"domain_name,omitempty"`

	// iam委托id
	IamAgencyId *string `json:"iam_agency_id,omitempty"`

	// iam委托名称
	IamAgencyName *string `json:"iam_agency_name,omitempty"`

	// 委托空间购买版本
	ResourceSpecCode *[]string `json:"resource_spec_code,omitempty"`

	// 是否被视图选中
	Selected *bool `json:"selected,omitempty"`
}

func (o CreateWorkspaceResponseBodyWorkspaceAgencyList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkspaceResponseBodyWorkspaceAgencyList struct{}"
	}

	return strings.Join([]string{"CreateWorkspaceResponseBodyWorkspaceAgencyList", string(data)}, " ")
}

type CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution struct {
	value string
}

type CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttributionEnum struct {
	THIS_ACCOUNT  CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution
	CROSS_ACCOUNT CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution
}

func GetCreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttributionEnum() CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttributionEnum {
	return CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttributionEnum{
		THIS_ACCOUNT: CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution{
			value: "THIS_ACCOUNT",
		},
		CROSS_ACCOUNT: CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution{
			value: "CROSS_ACCOUNT",
		},
	}
}

func (c CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution) Value() string {
	return c.value
}

func (c CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateWorkspaceResponseBodyWorkspaceAgencyListWorkspaceAttribution) UnmarshalJSON(b []byte) error {
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

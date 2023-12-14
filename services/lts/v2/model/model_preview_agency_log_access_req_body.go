package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PreviewAgencyLogAccessReqBody struct {

	// 日志访问类型
	AgencyAccessType PreviewAgencyLogAccessReqBodyAgencyAccessType `json:"agency_access_type"`

	// 跨账号日志接入配置名称
	AgencyLogAccess string `json:"agency_log_access"`

	// 委托日志流名称
	LogAgencyStreamName string `json:"log_agencyStream_name"`

	// 委托日志流id
	LogAgencyStreamId string `json:"log_agencyStream_id"`

	// 委托日志组名称
	LogAgencyGroupName string `json:"log_agencyGroup_name"`

	// 委托日志组id
	LogAgencyGroupId string `json:"log_agencyGroup_id"`

	// 被委托日志流名称
	LogBeAgencystreamName string `json:"log_beAgencystream_name"`

	// 被委托日志流id
	LogBeAgencystreamId string `json:"log_beAgencystream_id"`

	// 被委托日志组名称
	LogBeAgencygroupName string `json:"log_beAgencygroup_name"`

	// 被委托日志组id
	LogBeAgencygroupId string `json:"log_beAgencygroup_id"`

	// 被委托项目id
	BeAgencyProjectId string `json:"be_agency_project_id"`

	// 委托项目id
	AgencyProjectId string `json:"agency_project_id"`

	// 委托账号名称
	AgencyDomainName string `json:"agency_domain_name"`

	// 委托名称
	AgencyName string `json:"agency_name"`
}

func (o PreviewAgencyLogAccessReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewAgencyLogAccessReqBody struct{}"
	}

	return strings.Join([]string{"PreviewAgencyLogAccessReqBody", string(data)}, " ")
}

type PreviewAgencyLogAccessReqBodyAgencyAccessType struct {
	value string
}

type PreviewAgencyLogAccessReqBodyAgencyAccessTypeEnum struct {
	AGENCYACCESS PreviewAgencyLogAccessReqBodyAgencyAccessType
}

func GetPreviewAgencyLogAccessReqBodyAgencyAccessTypeEnum() PreviewAgencyLogAccessReqBodyAgencyAccessTypeEnum {
	return PreviewAgencyLogAccessReqBodyAgencyAccessTypeEnum{
		AGENCYACCESS: PreviewAgencyLogAccessReqBodyAgencyAccessType{
			value: "AGENCYACCESS",
		},
	}
}

func (c PreviewAgencyLogAccessReqBodyAgencyAccessType) Value() string {
	return c.value
}

func (c PreviewAgencyLogAccessReqBodyAgencyAccessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewAgencyLogAccessReqBodyAgencyAccessType) UnmarshalJSON(b []byte) error {
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

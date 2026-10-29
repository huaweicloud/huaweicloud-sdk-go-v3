package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LtsAccessConfigInfoRespon200 struct {

	// 跨账号日志接入id
	AccessConfigId *string `json:"access_config_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 跨账号日志接入名称
	AccessConfigName *string `json:"access_config_name,omitempty"`

	// **参数解释：** 委托接入类型。 **约束限制：** 不涉及。 **取值范围：** - AGENCYACCESS：委托接入。
	AccessConfigType *LtsAccessConfigInfoRespon200AccessConfigType `json:"access_config_type,omitempty"`

	// 日志组ID
	GroupId *string `json:"group_id,omitempty"`

	// 日志组名称
	LogGroupName *string `json:"log_group_name,omitempty"`

	// 日志流ID
	LogStreamId *string `json:"log_stream_id,omitempty"`

	// 日志流名称
	LogStreamName *string `json:"log_stream_name,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	AgencyLogAccess *PreviewAgencyLogAccessReqBody `json:"agency_log_access,omitempty"`
}

func (o LtsAccessConfigInfoRespon200) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LtsAccessConfigInfoRespon200 struct{}"
	}

	return strings.Join([]string{"LtsAccessConfigInfoRespon200", string(data)}, " ")
}

type LtsAccessConfigInfoRespon200AccessConfigType struct {
	value string
}

type LtsAccessConfigInfoRespon200AccessConfigTypeEnum struct {
	AGENT               LtsAccessConfigInfoRespon200AccessConfigType
	CRD                 LtsAccessConfigInfoRespon200AccessConfigType
	AGENCYACCESS        LtsAccessConfigInfoRespon200AccessConfigType
	K8_S_CCE            LtsAccessConfigInfoRespon200AccessConfigType
	K8_S_CUSTOM         LtsAccessConfigInfoRespon200AccessConfigType
	CLOUD_STREAM_AGENCY LtsAccessConfigInfoRespon200AccessConfigType
}

func GetLtsAccessConfigInfoRespon200AccessConfigTypeEnum() LtsAccessConfigInfoRespon200AccessConfigTypeEnum {
	return LtsAccessConfigInfoRespon200AccessConfigTypeEnum{
		AGENT: LtsAccessConfigInfoRespon200AccessConfigType{
			value: "Agent",
		},
		CRD: LtsAccessConfigInfoRespon200AccessConfigType{
			value: "CRD",
		},
		AGENCYACCESS: LtsAccessConfigInfoRespon200AccessConfigType{
			value: "AGENCYACCESS",
		},
		K8_S_CCE: LtsAccessConfigInfoRespon200AccessConfigType{
			value: "K8S_CCE",
		},
		K8_S_CUSTOM: LtsAccessConfigInfoRespon200AccessConfigType{
			value: "K8S_CUSTOM",
		},
		CLOUD_STREAM_AGENCY: LtsAccessConfigInfoRespon200AccessConfigType{
			value: "CLOUD_STREAM_AGENCY",
		},
	}
}

func (c LtsAccessConfigInfoRespon200AccessConfigType) Value() string {
	return c.value
}

func (c LtsAccessConfigInfoRespon200AccessConfigType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LtsAccessConfigInfoRespon200AccessConfigType) UnmarshalJSON(b []byte) error {
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

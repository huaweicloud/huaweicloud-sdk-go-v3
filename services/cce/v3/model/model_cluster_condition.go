package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ClusterCondition 集群详细状态。
type ClusterCondition struct {

	// **参数解释**： 状态类型。 **约束限制**： 不涉及 **取值范围**： - \"AgencyAvailable\": CCE集群自定义委托的状态。 - \"ClusterCertificate\": CCE集群证书的状态。 - \"ClusterCustomCertificate\": CCE集群自有证书的状态。 - \"CertificateRotation\": CCE集群证书更新的状态。 - \"CustomCertificateRotation\": CCE集群自有证书更新的状态。 - \"OpenIDConnectProcessing\": CCE集群开启OIDC特性处理中状态。 - \"OpenIDConnectProcessSuccess\": CCE集群开启OIDC特性成功状态。 - \"OpenIDConnectProcessFailed\": CCE集群开启OIDC特性失败状态。  **默认取值**： 不涉及
	Type *ClusterConditionType `json:"type,omitempty"`

	// **参数解释**： Condition当前状态。 **约束限制**： 不涉及 **取值范围**： - \"True\" - \"False\"  **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 上次状态检查时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	LastProbeTime *string `json:"lastProbeTime,omitempty"`

	// **参数解释**： 上次状态变更时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	LastTransitTime *string `json:"lastTransitTime,omitempty"`

	// **参数解释**： 上次状态变更原因。 **约束限制**： 不涉及 **取值范围**： - type为ClusterCertificate、ClusterCustomCertificate时取值范围：   - CertificateValid：证书状态正常   - CertificateExpiringWithin180Days：证书有效期低于180天   - CertificateExpiringWithin30Days：证书有效期低于30天   - CertificateExpired：证书已过期 - type为CertificateRotation、CustomCertificateRotation时取值范围：   - RotationInProgress：更新中   - RotationSucceeded：更新成功   - RotationFailed：更新失败  **默认取值**： 不涉及
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： Condition详细描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Message *string `json:"message,omitempty"`
}

func (o ClusterCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterCondition struct{}"
	}

	return strings.Join([]string{"ClusterCondition", string(data)}, " ")
}

type ClusterConditionType struct {
	value string
}

type ClusterConditionTypeEnum struct {
	AGENCY_AVAILABLE                ClusterConditionType
	CLUSTER_CERTIFICATE             ClusterConditionType
	CLUSTER_CUSTOM_CERTIFICATE      ClusterConditionType
	CERTIFICATE_ROTATION            ClusterConditionType
	CUSTOM_CERTIFICATE_ROTATION     ClusterConditionType
	OPEN_ID_CONNECT_PROCESSING      ClusterConditionType
	OPEN_ID_CONNECT_PROCESS_SUCCESS ClusterConditionType
	OPEN_ID_CONNECT_PROCESS_FAILED  ClusterConditionType
}

func GetClusterConditionTypeEnum() ClusterConditionTypeEnum {
	return ClusterConditionTypeEnum{
		AGENCY_AVAILABLE: ClusterConditionType{
			value: "AgencyAvailable",
		},
		CLUSTER_CERTIFICATE: ClusterConditionType{
			value: "ClusterCertificate",
		},
		CLUSTER_CUSTOM_CERTIFICATE: ClusterConditionType{
			value: "ClusterCustomCertificate",
		},
		CERTIFICATE_ROTATION: ClusterConditionType{
			value: "CertificateRotation",
		},
		CUSTOM_CERTIFICATE_ROTATION: ClusterConditionType{
			value: "CustomCertificateRotation",
		},
		OPEN_ID_CONNECT_PROCESSING: ClusterConditionType{
			value: "OpenIDConnectProcessing",
		},
		OPEN_ID_CONNECT_PROCESS_SUCCESS: ClusterConditionType{
			value: "OpenIDConnectProcessSuccess",
		},
		OPEN_ID_CONNECT_PROCESS_FAILED: ClusterConditionType{
			value: "OpenIDConnectProcessFailed",
		},
	}
}

func (c ClusterConditionType) Value() string {
	return c.value
}

func (c ClusterConditionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ClusterConditionType) UnmarshalJSON(b []byte) error {
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

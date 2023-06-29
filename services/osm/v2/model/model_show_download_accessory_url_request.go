package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDownloadAccessoryUrlRequest Request Object
type ShowDownloadAccessoryUrlRequest struct {

	// 附件id
	AccessoryId string `json:"accessory_id"`

	// - DEFAULT:  - NO_RELATION:  - NOTIFICATION:  - INCIDENT:  - MESSAGE:  - INSPECTION:  - CONNECT:
	RelationType ShowDownloadAccessoryUrlRequestRelationType `json:"relation_type"`

	// 关联id
	RelationId *string `json:"relation_id,omitempty"`

	// 工单id
	IncidentId *string `json:"incident_id,omitempty"`
}

func (o ShowDownloadAccessoryUrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDownloadAccessoryUrlRequest struct{}"
	}

	return strings.Join([]string{"ShowDownloadAccessoryUrlRequest", string(data)}, " ")
}

type ShowDownloadAccessoryUrlRequestRelationType struct {
	value string
}

type ShowDownloadAccessoryUrlRequestRelationTypeEnum struct {
	DEFAULT      ShowDownloadAccessoryUrlRequestRelationType
	NO_RELATION  ShowDownloadAccessoryUrlRequestRelationType
	NOTIFICATION ShowDownloadAccessoryUrlRequestRelationType
	INCIDENT     ShowDownloadAccessoryUrlRequestRelationType
	MESSAGE      ShowDownloadAccessoryUrlRequestRelationType
	INSPECTION   ShowDownloadAccessoryUrlRequestRelationType
	CONNECT      ShowDownloadAccessoryUrlRequestRelationType
}

func GetShowDownloadAccessoryUrlRequestRelationTypeEnum() ShowDownloadAccessoryUrlRequestRelationTypeEnum {
	return ShowDownloadAccessoryUrlRequestRelationTypeEnum{
		DEFAULT: ShowDownloadAccessoryUrlRequestRelationType{
			value: "DEFAULT",
		},
		NO_RELATION: ShowDownloadAccessoryUrlRequestRelationType{
			value: "NO_RELATION",
		},
		NOTIFICATION: ShowDownloadAccessoryUrlRequestRelationType{
			value: "NOTIFICATION",
		},
		INCIDENT: ShowDownloadAccessoryUrlRequestRelationType{
			value: "INCIDENT",
		},
		MESSAGE: ShowDownloadAccessoryUrlRequestRelationType{
			value: "MESSAGE",
		},
		INSPECTION: ShowDownloadAccessoryUrlRequestRelationType{
			value: "INSPECTION",
		},
		CONNECT: ShowDownloadAccessoryUrlRequestRelationType{
			value: "CONNECT",
		},
	}
}

func (c ShowDownloadAccessoryUrlRequestRelationType) Value() string {
	return c.value
}

func (c ShowDownloadAccessoryUrlRequestRelationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDownloadAccessoryUrlRequestRelationType) UnmarshalJSON(b []byte) error {
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

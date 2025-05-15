package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteConnectionStateEnum 分支网络连接状态。 - AVAILABLE （可用） - CREATING （创建中） - UPDATING （更新中） - DELETING （删除中） - FREEZING （冻结中） - UNFREEZING （解冻中） - RECOVERING （恢复中） - FAILED （失败） - FREEZED （冻结） - DELETED （已刪除）
type SiteConnectionStateEnum struct {
	value string
}

type SiteConnectionStateEnumEnum struct {
	AVAILABLE  SiteConnectionStateEnum
	CREATING   SiteConnectionStateEnum
	UPDATING   SiteConnectionStateEnum
	DELETING   SiteConnectionStateEnum
	FREEZING   SiteConnectionStateEnum
	UNFREEZING SiteConnectionStateEnum
	RECOVERING SiteConnectionStateEnum
	FAILED     SiteConnectionStateEnum
	FREEZED    SiteConnectionStateEnum
	DELETED    SiteConnectionStateEnum
}

func GetSiteConnectionStateEnumEnum() SiteConnectionStateEnumEnum {
	return SiteConnectionStateEnumEnum{
		AVAILABLE: SiteConnectionStateEnum{
			value: "AVAILABLE",
		},
		CREATING: SiteConnectionStateEnum{
			value: "CREATING",
		},
		UPDATING: SiteConnectionStateEnum{
			value: "UPDATING",
		},
		DELETING: SiteConnectionStateEnum{
			value: "DELETING",
		},
		FREEZING: SiteConnectionStateEnum{
			value: "FREEZING",
		},
		UNFREEZING: SiteConnectionStateEnum{
			value: "UNFREEZING",
		},
		RECOVERING: SiteConnectionStateEnum{
			value: "RECOVERING",
		},
		FAILED: SiteConnectionStateEnum{
			value: "FAILED",
		},
		FREEZED: SiteConnectionStateEnum{
			value: "FREEZED",
		},
		DELETED: SiteConnectionStateEnum{
			value: "DELETED",
		},
	}
}

func (c SiteConnectionStateEnum) Value() string {
	return c.value
}

func (c SiteConnectionStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteConnectionStateEnum) UnmarshalJSON(b []byte) error {
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

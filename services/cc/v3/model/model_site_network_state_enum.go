package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SiteNetworkStateEnum 分支网络状态。 - AVAILABLE (可用) - UPDATING (处理中) - FAILED (失败) - CREATING (创建中 - DELETING (删除中) - DELETED (已刪除) - NON-COMPLETE (配置未完成) - RESTORING (恢复中)
type SiteNetworkStateEnum struct {
	value string
}

type SiteNetworkStateEnumEnum struct {
	AVAILABLE    SiteNetworkStateEnum
	UPDATING     SiteNetworkStateEnum
	FAILED       SiteNetworkStateEnum
	CREATING     SiteNetworkStateEnum
	DELETING     SiteNetworkStateEnum
	DELETED      SiteNetworkStateEnum
	NON_COMPLETE SiteNetworkStateEnum
	RESTORING    SiteNetworkStateEnum
}

func GetSiteNetworkStateEnumEnum() SiteNetworkStateEnumEnum {
	return SiteNetworkStateEnumEnum{
		AVAILABLE: SiteNetworkStateEnum{
			value: "AVAILABLE",
		},
		UPDATING: SiteNetworkStateEnum{
			value: "UPDATING",
		},
		FAILED: SiteNetworkStateEnum{
			value: "FAILED",
		},
		CREATING: SiteNetworkStateEnum{
			value: "CREATING",
		},
		DELETING: SiteNetworkStateEnum{
			value: "DELETING",
		},
		DELETED: SiteNetworkStateEnum{
			value: "DELETED",
		},
		NON_COMPLETE: SiteNetworkStateEnum{
			value: "NON-COMPLETE",
		},
		RESTORING: SiteNetworkStateEnum{
			value: "RESTORING",
		},
	}
}

func (c SiteNetworkStateEnum) Value() string {
	return c.value
}

func (c SiteNetworkStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SiteNetworkStateEnum) UnmarshalJSON(b []byte) error {
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

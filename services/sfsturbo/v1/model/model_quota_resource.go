package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// QuotaResource 租户配额详情。
type QuotaResource struct {

	// 租户资源类型。shares表示文件系统数量，capacity表示文件系统容量。
	Type *QuotaResourceType `json:"type,omitempty"`

	// 已用配额。
	Used *int32 `json:"used,omitempty"`

	// 总配额。
	Quota *int32 `json:"quota,omitempty"`

	// 配额单位。
	Unit *string `json:"unit,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}

type QuotaResourceType struct {
	value string
}

type QuotaResourceTypeEnum struct {
	SHARES   QuotaResourceType
	CAPACITY QuotaResourceType
}

func GetQuotaResourceTypeEnum() QuotaResourceTypeEnum {
	return QuotaResourceTypeEnum{
		SHARES: QuotaResourceType{
			value: "shares",
		},
		CAPACITY: QuotaResourceType{
			value: "capacity",
		},
	}
}

func (c QuotaResourceType) Value() string {
	return c.value
}

func (c QuotaResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotaResourceType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AssetSharedConfig 共享配置
type AssetSharedConfig struct {

	// 共享类型。 * PRIVATE: 私有，仅本租户可访问。 * PUBLIC: 公开，所有租户可访问。当前仅提供系统资产可公开访问。 * SHARED：共享，指定租户可访问。拥有者指定租户可访问。
	SharedType *AssetSharedConfigSharedType `json:"shared_type,omitempty"`

	// 共享过期时间。默认过期时间为30天，即共享当天+30的23:59:59。
	ExpireTime *string `json:"expire_time,omitempty"`

	// 允许访问本资产的租户列表。
	AllowedProjectIds *[]string `json:"allowed_project_ids,omitempty"`
}

func (o AssetSharedConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssetSharedConfig struct{}"
	}

	return strings.Join([]string{"AssetSharedConfig", string(data)}, " ")
}

type AssetSharedConfigSharedType struct {
	value string
}

type AssetSharedConfigSharedTypeEnum struct {
	PRIVATE AssetSharedConfigSharedType
	PUBLIC  AssetSharedConfigSharedType
	SHARED  AssetSharedConfigSharedType
}

func GetAssetSharedConfigSharedTypeEnum() AssetSharedConfigSharedTypeEnum {
	return AssetSharedConfigSharedTypeEnum{
		PRIVATE: AssetSharedConfigSharedType{
			value: "PRIVATE",
		},
		PUBLIC: AssetSharedConfigSharedType{
			value: "PUBLIC",
		},
		SHARED: AssetSharedConfigSharedType{
			value: "SHARED",
		},
	}
}

func (c AssetSharedConfigSharedType) Value() string {
	return c.value
}

func (c AssetSharedConfigSharedType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AssetSharedConfigSharedType) UnmarshalJSON(b []byte) error {
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

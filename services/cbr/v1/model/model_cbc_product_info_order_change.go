package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CbcProductInfoOrderChange struct {

	// 产品标识，通过订购询价接口获得，长度限制：1-64，只能由字母、数字、“_”、“-”组成。
	ProductId string `json:"product_id"`

	// 资源容量大小，取值范围：10-10485760
	ResourceSize int32 `json:"resource_size"`

	// 资源容量度量标识，枚举值17：GB
	ResourceSizeMeasureId *int32 `json:"resource_size_measure_id,omitempty"`

	// 用户购买云服务产品的资源规格 Enum: [vault.backup.server.normal，vault.backup.turbo.normal, vault.backup.database.normal，vault.backup.volume.normal，vault.backup.rds.normal，vault.replication.server.normal，vault.hybrid.server.normal]
	ResourceSpecCode CbcProductInfoOrderChangeResourceSpecCode `json:"resource_spec_code"`
}

func (o CbcProductInfoOrderChange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcProductInfoOrderChange struct{}"
	}

	return strings.Join([]string{"CbcProductInfoOrderChange", string(data)}, " ")
}

type CbcProductInfoOrderChangeResourceSpecCode struct {
	value string
}

type CbcProductInfoOrderChangeResourceSpecCodeEnum struct {
	VAULT_BACKUP_SERVER_NORMAL      CbcProductInfoOrderChangeResourceSpecCode
	VAULT_BACKUP_TURBO_NORMAL       CbcProductInfoOrderChangeResourceSpecCode
	VAULT_BACKUP_DATABASE_NORMAL    CbcProductInfoOrderChangeResourceSpecCode
	VAULT_BACKUP_VOLUME_NORMAL      CbcProductInfoOrderChangeResourceSpecCode
	VAULT_BACKUP_RDS_NORMAL         CbcProductInfoOrderChangeResourceSpecCode
	VAULT_REPLICATION_SERVER_NORMAL CbcProductInfoOrderChangeResourceSpecCode
	VAULT_HYBRID_SERVER_NORMAL      CbcProductInfoOrderChangeResourceSpecCode
}

func GetCbcProductInfoOrderChangeResourceSpecCodeEnum() CbcProductInfoOrderChangeResourceSpecCodeEnum {
	return CbcProductInfoOrderChangeResourceSpecCodeEnum{
		VAULT_BACKUP_SERVER_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.backup.server.normal",
		},
		VAULT_BACKUP_TURBO_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.backup.turbo.normal",
		},
		VAULT_BACKUP_DATABASE_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.backup.database.normal",
		},
		VAULT_BACKUP_VOLUME_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.backup.volume.normal",
		},
		VAULT_BACKUP_RDS_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.backup.rds.normal",
		},
		VAULT_REPLICATION_SERVER_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.replication.server.normal",
		},
		VAULT_HYBRID_SERVER_NORMAL: CbcProductInfoOrderChangeResourceSpecCode{
			value: "vault.hybrid.server.normal",
		},
	}
}

func (c CbcProductInfoOrderChangeResourceSpecCode) Value() string {
	return c.value
}

func (c CbcProductInfoOrderChangeResourceSpecCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CbcProductInfoOrderChangeResourceSpecCode) UnmarshalJSON(b []byte) error {
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

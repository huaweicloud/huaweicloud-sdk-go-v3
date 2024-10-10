package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StarRocksCreateRequest struct {

	// 实例名称。同一租户下，同类型的实例名可重名。  取值范围：最小为4个字符，最大为64个字符且不超过64个字节，必须以字母开头，区分大小写，可以包含字母、数字、中划线、下划线，不能包含其他特殊字符。不支持中文名。
	Name string `json:"name"`

	Engine *StarRocksCreateRequestEngine `json:"engine"`

	Ha *StarRocksCreateRequestHa `json:"ha"`

	// FE节点规格ID。使用可通过查询HTAP规格响应消息中的“id”。
	FeFlavorId string `json:"fe_flavor_id"`

	// BE节点规格ID。使用可通过查询HTAP规格响应消息中的“id”。
	BeFlavorId string `json:"be_flavor_id"`

	// 数据库密码。  取值范围：至少包含以下字符的三种：大小写字母、数字和特殊符号~!@#$%^*-_=+?,()&|.，长度8~32个字符。 建议您输入高强度密码，以提高安全性，防止出现密码被暴力破解等安全风险。如果您输入弱密码，系统会自动判定密码非法。
	DbRootPwd string `json:"db_root_pwd"`

	// FE节点数。 - 单机时固定为1 - 集群时取值[3, 10]
	FeCount int32 `json:"fe_count"`

	// BE节点数。 - 单机时固定为1 - 集群时取值[3, 10]
	BeCount int32 `json:"be_count"`

	// 可用区类型。 当前仅支持single。
	AzMode StarRocksCreateRequestAzMode `json:"az_mode"`

	FeVolume *StarRocksCreateRequestFeVolume `json:"fe_volume"`

	BeVolume *StarRocksCreateRequestBeVolume `json:"be_volume"`

	// 可用区代码。
	AzCode string `json:"az_code"`

	// 时区。默认时区为UTC+08:00。
	TimeZone *string `json:"time_zone,omitempty"`

	TagsInfo *StarRocksCreateRequestTagsInfo `json:"tags_info"`

	// 实例安全组ID。默认与Taurus安全组ID一致。
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o StarRocksCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StarRocksCreateRequest struct{}"
	}

	return strings.Join([]string{"StarRocksCreateRequest", string(data)}, " ")
}

type StarRocksCreateRequestAzMode struct {
	value string
}

type StarRocksCreateRequestAzModeEnum struct {
	SINGLE StarRocksCreateRequestAzMode
}

func GetStarRocksCreateRequestAzModeEnum() StarRocksCreateRequestAzModeEnum {
	return StarRocksCreateRequestAzModeEnum{
		SINGLE: StarRocksCreateRequestAzMode{
			value: "single",
		},
	}
}

func (c StarRocksCreateRequestAzMode) Value() string {
	return c.value
}

func (c StarRocksCreateRequestAzMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StarRocksCreateRequestAzMode) UnmarshalJSON(b []byte) error {
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

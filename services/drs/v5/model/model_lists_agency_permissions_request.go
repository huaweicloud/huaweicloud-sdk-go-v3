package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListsAgencyPermissionsRequest Request Object
type ListsAgencyPermissionsRequest struct {

	// 请求语言类型。
	XLanguage *ListsAgencyPermissionsRequestXLanguage `json:"X-Language,omitempty"`

	// 源库类型 - mysql - sqlserver - postgresql - hwsql - mongodb - dws - oracle - taurus - tauruslite - ddm - kafka - mrsKafka - gaussdb - gaussdbv5 - gaussdbv5ha - gaussmongodb - cassandra - dmq - gaussdbt - gaussdb300 - gaussdbtha - elasticsearch - db2 - tidb - redis - rediscluster - gaussredis - mariadb - gaussdbv1 - informix - dynamo - gausscassandra - oceanbase
	SourceType *string `json:"source_type,omitempty"`

	// 目标库类型 - mysql - sqlserver - postgresql - hwsql - mongodb - dws - oracle - taurus - tauruslite - ddm - kafka - mrsKafka - gaussdb - gaussdbv5 - gaussdbv5ha - gaussmongodb - cassandra - dmq - gaussdbt - gaussdb300 - gaussdbtha - elasticsearch - db2 - tidb - redis - rediscluster - gaussredis - mariadb - gaussdbv1 - informix - dynamo - gausscassandra - oceanbase
	TargetType *string `json:"target_type,omitempty"`

	// 是否自建。
	IsNonDbs bool `json:"is_non_dbs"`
}

func (o ListsAgencyPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListsAgencyPermissionsRequest struct{}"
	}

	return strings.Join([]string{"ListsAgencyPermissionsRequest", string(data)}, " ")
}

type ListsAgencyPermissionsRequestXLanguage struct {
	value string
}

type ListsAgencyPermissionsRequestXLanguageEnum struct {
	EN_US ListsAgencyPermissionsRequestXLanguage
	ZH_CN ListsAgencyPermissionsRequestXLanguage
}

func GetListsAgencyPermissionsRequestXLanguageEnum() ListsAgencyPermissionsRequestXLanguageEnum {
	return ListsAgencyPermissionsRequestXLanguageEnum{
		EN_US: ListsAgencyPermissionsRequestXLanguage{
			value: "en-us",
		},
		ZH_CN: ListsAgencyPermissionsRequestXLanguage{
			value: "zh-cn",
		},
	}
}

func (c ListsAgencyPermissionsRequestXLanguage) Value() string {
	return c.value
}

func (c ListsAgencyPermissionsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListsAgencyPermissionsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

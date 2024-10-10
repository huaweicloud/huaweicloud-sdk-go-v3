package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// 请求语言类型。
	XLanguage *string `json:"X-Language,omitempty"`

	// 连接ID。
	ConnectionId *string `json:"connection_id,omitempty"`

	// 连接类型。 - mysql - oracle - postgresql - mongodb
	DbType *ListConnectionsRequestDbType `json:"db_type,omitempty"`

	// 连接名称，忽略大小写。
	Name *string `json:"name,omitempty"`

	// 云上数据库实例ID。
	InstId *string `json:"inst_id,omitempty"`

	// 连接IP。
	Ip *string `json:"ip,omitempty"`

	// 连接描述。
	Description *string `json:"description,omitempty"`

	// 时间区间用“，”分割。示例：2024-05-17T07:46:00.414Z,2024-05-20T07:46:00.999Z。
	CreateTime *string `json:"create_time,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量，默认值为0，表示查询该偏移量后面的记录。
	Offset *int32 `json:"offset,omitempty"`

	// 查询返回记录的数量限制，默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 值为“true”时会使得offset和limit参数失效并返回所有记录。
	FetchAll *bool `json:"fetch_all,omitempty"`

	// 返回结果按该关键字排序，默认为“created_at”。
	SortKey *string `json:"sort_key,omitempty"`

	// 降序或升序（分别对应desc和asc，默认为“desc”）。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}

type ListConnectionsRequestDbType struct {
	value string
}

type ListConnectionsRequestDbTypeEnum struct {
	MYSQL      ListConnectionsRequestDbType
	ORACLE     ListConnectionsRequestDbType
	POSTGRESQL ListConnectionsRequestDbType
	MONGODB    ListConnectionsRequestDbType
}

func GetListConnectionsRequestDbTypeEnum() ListConnectionsRequestDbTypeEnum {
	return ListConnectionsRequestDbTypeEnum{
		MYSQL: ListConnectionsRequestDbType{
			value: "mysql",
		},
		ORACLE: ListConnectionsRequestDbType{
			value: "oracle",
		},
		POSTGRESQL: ListConnectionsRequestDbType{
			value: "postgresql",
		},
		MONGODB: ListConnectionsRequestDbType{
			value: "mongodb",
		},
	}
}

func (c ListConnectionsRequestDbType) Value() string {
	return c.value
}

func (c ListConnectionsRequestDbType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListConnectionsRequestDbType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListTaskTableReferenceDetailRequest Request Object
type ListTaskTableReferenceDetailRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 表名。
	TableName string `json:"table_name"`

	// 数据库类型，仅支持DLI，HIVE，SPARK。
	DbType string `json:"db_type"`

	// 数据库名称。
	DataBaseName *string `json:"data_base_name,omitempty"`

	// 集群名称。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 输入输出类型： - 0: 读表 - 1: 写表
	IoType *ListTaskTableReferenceDetailRequestIoType `json:"io_type,omitempty"`

	// 分页的起始页，取值范围大于等于0。默认值: 0。
	Offset *int32 `json:"offset,omitempty"`

	// 分页返回结果，指定每页最大记录数。默认值: 20。
	Limit *int32 `json:"limit,omitempty"`

	// 工作空间名称。
	WorkspaceName *string `json:"workspace_name,omitempty"`

	// 作业责任人。
	Owner *string `json:"owner,omitempty"`

	// 作业执行用户。
	ExecuteUser *string `json:"execute_user,omitempty"`
}

func (o ListTaskTableReferenceDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskTableReferenceDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTaskTableReferenceDetailRequest", string(data)}, " ")
}

type ListTaskTableReferenceDetailRequestIoType struct {
	value int32
}

type ListTaskTableReferenceDetailRequestIoTypeEnum struct {
	E_0 ListTaskTableReferenceDetailRequestIoType
	E_1 ListTaskTableReferenceDetailRequestIoType
}

func GetListTaskTableReferenceDetailRequestIoTypeEnum() ListTaskTableReferenceDetailRequestIoTypeEnum {
	return ListTaskTableReferenceDetailRequestIoTypeEnum{
		E_0: ListTaskTableReferenceDetailRequestIoType{
			value: 0,
		}, E_1: ListTaskTableReferenceDetailRequestIoType{
			value: 1,
		},
	}
}

func (c ListTaskTableReferenceDetailRequestIoType) Value() int32 {
	return c.value
}

func (c ListTaskTableReferenceDetailRequestIoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListTaskTableReferenceDetailRequestIoType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

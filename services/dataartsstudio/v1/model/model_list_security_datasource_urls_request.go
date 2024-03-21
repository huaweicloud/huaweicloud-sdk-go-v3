package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSecurityDatasourceUrlsRequest Request Object
type ListSecurityDatasourceUrlsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 数据源类型,HIVE
	DatasourceType *ListSecurityDatasourceUrlsRequestDatasourceType `json:"datasource_type,omitempty"`

	// 父权限集ID。获取方法请参见[查询权限集列表](ListSecurityPermissionSets.xml) 注意： * 当该值为父权限集ID时，则基于父权限集中的权限查询
	ParentPermissionSetId *string `json:"parent_permission_set_id,omitempty"`
}

func (o ListSecurityDatasourceUrlsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityDatasourceUrlsRequest struct{}"
	}

	return strings.Join([]string{"ListSecurityDatasourceUrlsRequest", string(data)}, " ")
}

type ListSecurityDatasourceUrlsRequestDatasourceType struct {
	value string
}

type ListSecurityDatasourceUrlsRequestDatasourceTypeEnum struct {
	HIVE ListSecurityDatasourceUrlsRequestDatasourceType
}

func GetListSecurityDatasourceUrlsRequestDatasourceTypeEnum() ListSecurityDatasourceUrlsRequestDatasourceTypeEnum {
	return ListSecurityDatasourceUrlsRequestDatasourceTypeEnum{
		HIVE: ListSecurityDatasourceUrlsRequestDatasourceType{
			value: "HIVE",
		},
	}
}

func (c ListSecurityDatasourceUrlsRequestDatasourceType) Value() string {
	return c.value
}

func (c ListSecurityDatasourceUrlsRequestDatasourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSecurityDatasourceUrlsRequestDatasourceType) UnmarshalJSON(b []byte) error {
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

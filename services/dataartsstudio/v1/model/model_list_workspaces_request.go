package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkspacesRequest Request Object
type ListWorkspacesRequest struct {

	// 工作空间ID，获取方法请参见[实例ID和工作空间ID](dataartsstudio_02_0350.xml)。
	Workspace string `json:"workspace"`

	// 项目ID，获取方法请参见[项目ID和账号ID](projectid_accountid.xml)。  多project场景采用AK/SK认证的接口请求，则该字段必选。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 默认值：application/json;charset=UTF-8 可选，有Body体的情况下必选，没有Body体则无需填写和校验。
	ContentType *string `json:"Content-Type,omitempty"`

	// 模型工作区类型枚举。 枚举值：   - THIRD_NF: 关系建模   - DIMENSION: 维度建模
	WorkspaceType *ListWorkspacesRequestWorkspaceType `json:"workspace_type,omitempty"`

	// 每页查询条数，即查询Y条数据。默认值50，取值范围[1,100]。
	Limit *int32 `json:"limit,omitempty"`

	// 查询起始坐标，即跳过X条数据，仅支持0或limit的整数倍，不满足则向下取整，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`
}

func (o ListWorkspacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkspacesRequest struct{}"
	}

	return strings.Join([]string{"ListWorkspacesRequest", string(data)}, " ")
}

type ListWorkspacesRequestWorkspaceType struct {
	value string
}

type ListWorkspacesRequestWorkspaceTypeEnum struct {
	THIRD_NF  ListWorkspacesRequestWorkspaceType
	DIMENSION ListWorkspacesRequestWorkspaceType
}

func GetListWorkspacesRequestWorkspaceTypeEnum() ListWorkspacesRequestWorkspaceTypeEnum {
	return ListWorkspacesRequestWorkspaceTypeEnum{
		THIRD_NF: ListWorkspacesRequestWorkspaceType{
			value: "THIRD_NF",
		},
		DIMENSION: ListWorkspacesRequestWorkspaceType{
			value: "DIMENSION",
		},
	}
}

func (c ListWorkspacesRequestWorkspaceType) Value() string {
	return c.value
}

func (c ListWorkspacesRequestWorkspaceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkspacesRequestWorkspaceType) UnmarshalJSON(b []byte) error {
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

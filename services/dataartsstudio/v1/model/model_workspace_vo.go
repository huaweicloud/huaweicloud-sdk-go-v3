package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type WorkspaceVo struct {

	// 编号
	Id *int64 `json:"id,omitempty"`

	// 工作区名字
	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	// 是否为物理表
	IsPhysical *bool `json:"is_physical,omitempty"`

	// 是否为常用
	Frequent *bool `json:"frequent,omitempty"`

	// 分层治理
	Top *bool `json:"top,omitempty"`

	Level *ModelLevel `json:"level,omitempty"`

	// 数据连接类型
	DwType *string `json:"dw_type,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 工作区类型枚举
	Type WorkspaceVoType `json:"type"`

	// 关联的业务分层的id列表 {\"l1Ids\":[],\"l2Ids\":[],\"l3Ids\":[]}
	BizCatalogIds *string `json:"biz_catalog_ids,omitempty"`

	// 数据库名称数组
	Databases *[]string `json:"databases,omitempty"`
}

func (o WorkspaceVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkspaceVo struct{}"
	}

	return strings.Join([]string{"WorkspaceVo", string(data)}, " ")
}

type WorkspaceVoType struct {
	value string
}

type WorkspaceVoTypeEnum struct {
	THIRD_NF  WorkspaceVoType
	DIMENSION WorkspaceVoType
}

func GetWorkspaceVoTypeEnum() WorkspaceVoTypeEnum {
	return WorkspaceVoTypeEnum{
		THIRD_NF: WorkspaceVoType{
			value: "THIRD_NF",
		},
		DIMENSION: WorkspaceVoType{
			value: "DIMENSION",
		},
	}
}

func (c WorkspaceVoType) Value() string {
	return c.value
}

func (c WorkspaceVoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkspaceVoType) UnmarshalJSON(b []byte) error {
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

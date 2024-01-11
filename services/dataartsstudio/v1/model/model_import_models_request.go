package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImportModelsRequest Request Object
type ImportModelsRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 需要执行的动作，根据导入的对象不同而选择不同的导入动作
	ActionId ImportModelsRequestActionId `json:"action-id"`

	// 模型id，在导入模型（import_relation）时必填
	ModelId *string `json:"model_id,omitempty"`

	// 导入的目录id，在导入码表（import_datastandard）和数据标准（import_datastandard）时生效，选填
	DirectoryId *string `json:"directory_id,omitempty"`

	// 是否需要覆盖更新已有的实体
	SkipExist *bool `json:"skip-exist,omitempty"`

	Body *ImportModelsRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o ImportModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportModelsRequest struct{}"
	}

	return strings.Join([]string{"ImportModelsRequest", string(data)}, " ")
}

type ImportModelsRequestActionId struct {
	value string
}

type ImportModelsRequestActionIdEnum struct {
	IMPORT_RELATION     ImportModelsRequestActionId
	IMPORT_DIMENSION    ImportModelsRequestActionId
	IMPORT_CODETABLE    ImportModelsRequestActionId
	IMPORT_DATASTANDARD ImportModelsRequestActionId
	IMPORT_BIZMETRIC    ImportModelsRequestActionId
	IMPORT_BIZCATALOG   ImportModelsRequestActionId
	IMPORT_ATOMIC       ImportModelsRequestActionId
	IMPORT_DERIVATIVE   ImportModelsRequestActionId
	IMPORT_COMPOUND     ImportModelsRequestActionId
	IMPORT_AGGREGATION  ImportModelsRequestActionId
}

func GetImportModelsRequestActionIdEnum() ImportModelsRequestActionIdEnum {
	return ImportModelsRequestActionIdEnum{
		IMPORT_RELATION: ImportModelsRequestActionId{
			value: "import_relation",
		},
		IMPORT_DIMENSION: ImportModelsRequestActionId{
			value: "import_dimension",
		},
		IMPORT_CODETABLE: ImportModelsRequestActionId{
			value: "import_codetable",
		},
		IMPORT_DATASTANDARD: ImportModelsRequestActionId{
			value: "import_datastandard",
		},
		IMPORT_BIZMETRIC: ImportModelsRequestActionId{
			value: "import_bizmetric",
		},
		IMPORT_BIZCATALOG: ImportModelsRequestActionId{
			value: "import_bizcatalog",
		},
		IMPORT_ATOMIC: ImportModelsRequestActionId{
			value: "import_atomic",
		},
		IMPORT_DERIVATIVE: ImportModelsRequestActionId{
			value: "import_derivative",
		},
		IMPORT_COMPOUND: ImportModelsRequestActionId{
			value: "import_compound",
		},
		IMPORT_AGGREGATION: ImportModelsRequestActionId{
			value: "import_aggregation",
		},
	}
}

func (c ImportModelsRequestActionId) Value() string {
	return c.value
}

func (c ImportModelsRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportModelsRequestActionId) UnmarshalJSON(b []byte) error {
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

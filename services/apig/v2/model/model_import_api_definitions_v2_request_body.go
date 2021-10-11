package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"strings"
)

type ImportApiDefinitionsV2RequestBody struct {
	// 是否创建新分组
	IsCreateGroup *def.MultiPart `json:"is_create_group,omitempty"`

	// API分组编号，当is_create_group=false时为必填
	GroupId *def.MultiPart `json:"group_id,omitempty"`

	// 扩展信息导入模式 - merge：当扩展信息定义冲突时，merge保留原有扩展信息 - override：当扩展信息定义冲突时，override会覆盖原有扩展信息
	ExtendMode *def.MultiPart `json:"extend_mode,omitempty"`

	// 是否开启简易导入模式
	SimpleMode *def.MultiPart `json:"simple_mode,omitempty"`

	// 是否开启Mock后端
	MockMode *def.MultiPart `json:"mock_mode,omitempty"`

	// 导入模式 - merge：当API信息定义冲突时，merge保留原有API信息 - override：当API信息定义冲突时，override会覆盖原有API信息
	ApiMode *def.MultiPart `json:"api_mode,omitempty"`

	// 导入Api的请求体，json或yaml格式的文件
	FileName *def.FilePart `json:"file_name"`
}

func (o ImportApiDefinitionsV2RequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportApiDefinitionsV2RequestBody struct{}"
	}

	return strings.Join([]string{"ImportApiDefinitionsV2RequestBody", string(data)}, " ")
}

type ImportApiDefinitionsV2RequestBodyExtendMode struct {
	value string
}

type ImportApiDefinitionsV2RequestBodyExtendModeEnum struct {
	MERGE    ImportApiDefinitionsV2RequestBodyExtendMode
	OVERRIDE ImportApiDefinitionsV2RequestBodyExtendMode
}

func GetImportApiDefinitionsV2RequestBodyExtendModeEnum() ImportApiDefinitionsV2RequestBodyExtendModeEnum {
	return ImportApiDefinitionsV2RequestBodyExtendModeEnum{
		MERGE: ImportApiDefinitionsV2RequestBodyExtendMode{
			value: "merge",
		},
		OVERRIDE: ImportApiDefinitionsV2RequestBodyExtendMode{
			value: "override",
		},
	}
}

func (c ImportApiDefinitionsV2RequestBodyExtendMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ImportApiDefinitionsV2RequestBodyExtendMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ImportApiDefinitionsV2RequestBodyApiMode struct {
	value string
}

type ImportApiDefinitionsV2RequestBodyApiModeEnum struct {
	MERGE    ImportApiDefinitionsV2RequestBodyApiMode
	OVERRIDE ImportApiDefinitionsV2RequestBodyApiMode
}

func GetImportApiDefinitionsV2RequestBodyApiModeEnum() ImportApiDefinitionsV2RequestBodyApiModeEnum {
	return ImportApiDefinitionsV2RequestBodyApiModeEnum{
		MERGE: ImportApiDefinitionsV2RequestBodyApiMode{
			value: "merge",
		},
		OVERRIDE: ImportApiDefinitionsV2RequestBodyApiMode{
			value: "override",
		},
	}
}

func (c ImportApiDefinitionsV2RequestBodyApiMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ImportApiDefinitionsV2RequestBodyApiMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHpcCacheTaskReq hpc cache 联动任务请求
type CreateHpcCacheTaskReq struct {

	// 任务类型，当前支持import(附加元数据导入)，import_metadata(快速导入)，preload(数据预热)，export(导出)。  附加元数据导入方式会导入OBS对象的元数据（名称、大小、最后修改时间）以及来源于SFS Turbo 导出时的附加元数据（如uid、gid、mode）。  快速导入方式仅会导入OBS对象的元数据（名称、大小、最后修改时间），不会导入其它附加元数据（如uid、gid、mode），SFS Turbo会生成默认的附加元数据。  数据预热功能会同时导入元数据和数据内容，数据预热中的元数据导入采用快速导入方式，不会导入其它附加元数据（如uid、gid、mode）。  数据导出功能会将您在联动目录里创建的文件，以及对从OBS导入后又做过修改的文件导出存储到OBS桶里。
	Type CreateHpcCacheTaskReqType `json:"type"`

	// 联动目录名称
	SrcTarget string `json:"src_target"`

	// 导入导出任务的源端路径前缀，导入时不需要包含OBS桶名，导出时不需要包含联动目录名称。  当 type 为 import、import_metadata、preload 时，该参数必须是 OBS后端存储内的相对路径前缀。 当 type 为 export 时，该参数必须是文件系统内file_system_path下的相对路径前缀。  如果不带该字段，导入时会导入绑定OBS桶内的所有对象，导出时会导出联动目录下的所有文件。限制如下： - 长度为1~1024个字符之间 - 不能包含以下字符：\\:*?\"<>| - 不能以英文句号.开头和结尾
	SrcPrefix *string `json:"src_prefix,omitempty"`

	// 目前只支持和src_target保持一致
	DestTarget string `json:"dest_target"`

	// 目前只支持和src_prefix保持一致
	DestPrefix *string `json:"dest_prefix,omitempty"`

	Attributes *ObsTargetAttributes `json:"attributes,omitempty"`
}

func (o CreateHpcCacheTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHpcCacheTaskReq struct{}"
	}

	return strings.Join([]string{"CreateHpcCacheTaskReq", string(data)}, " ")
}

type CreateHpcCacheTaskReqType struct {
	value string
}

type CreateHpcCacheTaskReqTypeEnum struct {
	IMPORT          CreateHpcCacheTaskReqType
	EXPORT          CreateHpcCacheTaskReqType
	IMPORT_METADATA CreateHpcCacheTaskReqType
	PRELOAD         CreateHpcCacheTaskReqType
}

func GetCreateHpcCacheTaskReqTypeEnum() CreateHpcCacheTaskReqTypeEnum {
	return CreateHpcCacheTaskReqTypeEnum{
		IMPORT: CreateHpcCacheTaskReqType{
			value: "import",
		},
		EXPORT: CreateHpcCacheTaskReqType{
			value: "export",
		},
		IMPORT_METADATA: CreateHpcCacheTaskReqType{
			value: "import_metadata",
		},
		PRELOAD: CreateHpcCacheTaskReqType{
			value: "preload",
		},
	}
}

func (c CreateHpcCacheTaskReqType) Value() string {
	return c.value
}

func (c CreateHpcCacheTaskReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHpcCacheTaskReqType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHpcCacheTaskReq hpc cache 联动任务请求
type CreateHpcCacheTaskReq struct {

	// 操作类型，当前支持import(普通导入)，export(导出)，import_metadata(元数据导入)，preload(数据预热)
	Type CreateHpcCacheTaskReqType `json:"type"`

	// 源端对象。OBS桶绑定文件系统后的文件系统路径名称
	SrcTarget string `json:"src_target"`

	// 源端路径前缀。例如，'/mnt/sfs_turbo' 为您的挂载目录，对于导入任务，前缀为prefix，则会导入到 '/mnt/sfs_turbo/prefix'；如导入前缀为空，则会直接导入到 '/mnt/sfs_turbo'。对于导出任务，前缀为prefix，则会导出到 '/mnt/sfs_turbo/prefix'；如导出前缀为空，则会直接导出到 '/mnt/sfs_turbo'。
	SrcPrefix *string `json:"src_prefix,omitempty"`

	// 目的端对象。目前只支持和src_target保持一致
	DestTarget string `json:"dest_target"`

	// 目的端路径。目前只支持和src_prefix保持一致
	DestPrefix *string `json:"dest_prefix,omitempty"`
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

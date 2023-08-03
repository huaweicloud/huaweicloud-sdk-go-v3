package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HumanModelMetaProperties struct {

	// 当前模型中的WHOLE_MODEL是基于哪个file_id生成的，如果当前记录的信息与MAIN文件的file_id一致，那就认为已经生成过，无需再进行全模型导出
	WholeModelBaseFileId *string `json:"whole_model_base_file_id,omitempty"`

	// 当前用于加载的file_id信息，若为空或未匹配到，则使用MAIN文件
	LoadModelFileId *string `json:"load_model_file_id,omitempty"`
}

func (o HumanModelMetaProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanModelMetaProperties struct{}"
	}

	return strings.Join([]string{"HumanModelMetaProperties", string(data)}, " ")
}

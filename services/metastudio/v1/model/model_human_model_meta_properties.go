package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HumanModelMetaProperties struct {

	// **参数解释**： 用于生成WHOLE_MODEL的模型file_id **约束限制**： 如果当前记录的信息与MAIN文件的file_id一致，那就认为已经生成过，无需再进行全模型导出 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及。
	WholeModelBaseFileId *string `json:"whole_model_base_file_id,omitempty"`

	// **参数解释**： 当前用于渲染加载的模型file_id **约束限制**： 若为空或未匹配到，则使用MAIN文件 **取值范围**： 字符长度0-64位。 **默认取值**： 不涉及。
	LoadModelFileId *string `json:"load_model_file_id,omitempty"`
}

func (o HumanModelMetaProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanModelMetaProperties struct{}"
	}

	return strings.Join([]string{"HumanModelMetaProperties", string(data)}, " ")
}

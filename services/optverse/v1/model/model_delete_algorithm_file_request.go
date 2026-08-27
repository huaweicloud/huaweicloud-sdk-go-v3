package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlgorithmFileRequest Request Object
type DeleteAlgorithmFileRequest struct {

	// **参数解释**： 算法项目标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	AlgorithmId string `json:"algorithm_id"`

	// **参数解释**： 算法层级路径。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-896]个字符。 **默认取值**： 不涉及
	FilePath string `json:"file_path"`
}

func (o DeleteAlgorithmFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlgorithmFileRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlgorithmFileRequest", string(data)}, " ")
}

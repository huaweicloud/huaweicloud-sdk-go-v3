package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveAlgorithmFileRequest Request Object
type SaveAlgorithmFileRequest struct {

	// **参数解释**： 算法项目标识符。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	AlgorithmId string `json:"algorithm_id"`

	Body *SaveAlgorithmFileRequestBody `json:"body,omitempty" type:"multipart"`
}

func (o SaveAlgorithmFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveAlgorithmFileRequest struct{}"
	}

	return strings.Join([]string{"SaveAlgorithmFileRequest", string(data)}, " ")
}

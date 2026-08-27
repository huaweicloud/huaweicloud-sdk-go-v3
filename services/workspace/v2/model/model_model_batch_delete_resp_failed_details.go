package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModelBatchDeleteRespFailedDetails struct {

	// 模型id。
	ModelId *string `json:"model_id,omitempty"`

	// 失败原因。
	Reason *string `json:"reason,omitempty"`
}

func (o ModelBatchDeleteRespFailedDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelBatchDeleteRespFailedDetails struct{}"
	}

	return strings.Join([]string{"ModelBatchDeleteRespFailedDetails", string(data)}, " ")
}

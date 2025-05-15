package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AutoIdDocClassificationResult struct {

	// 证件的类型，支持的证件类型如表1所示：   **表1* 支持的证件类型 | 证件类型               | 类型描述                 | | ---------------------- | ------------------------ | | peru_id_card           | 秘鲁身份证               | | cambodian_id_card      | 柬文身份证               | | hongkong_id_card       | 香港身份证               | | macao_id_card          | 澳门身份证               | | myanmar_driver_license | 缅文驾驶证               | | myanmar_id_card        | 缅文身份证               | | passport               | 护照                     | | thailand_id_card       | 泰文身份证               | | id_card                | 中华人民共和国居民身份证 |
	Type *string `json:"type,omitempty"`

	// 证件的位置。
	Location *[][]int32 `json:"location,omitempty"`

	// 证件位置的置信度。
	Confidence *float32 `json:"confidence,omitempty"`

	AlarmResult *AutoIdDocClassificationAlarmResult `json:"alarm_result,omitempty"`

	AlarmConfidence *AutoIdDocClassificationAlarmConfidence `json:"alarm_confidence,omitempty"`
}

func (o AutoIdDocClassificationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoIdDocClassificationResult struct{}"
	}

	return strings.Join([]string{"AutoIdDocClassificationResult", string(data)}, " ")
}

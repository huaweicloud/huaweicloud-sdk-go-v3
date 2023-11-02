package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest Request Object
type ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest struct {

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest struct{}"
	}

	return strings.Join([]string{"ShowIntelligentDiagnosisAbnormalCountOfInstancesRequest", string(data)}, " ")
}

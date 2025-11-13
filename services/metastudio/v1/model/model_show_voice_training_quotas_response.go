package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVoiceTrainingQuotasResponse Response Object
type ShowVoiceTrainingQuotasResponse struct {

	// 声音模型训练资源总条数。
	Count *int32 `json:"count,omitempty"`

	ResourceAvailableNums *ResourceAvailableNums `json:"resource_available_nums,omitempty"`

	// 声音模型训练资源列表。
	Resources      *[]VoiceTrainingResource `json:"resources,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowVoiceTrainingQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVoiceTrainingQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowVoiceTrainingQuotasResponse", string(data)}, " ")
}

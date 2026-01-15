package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAiOpsDetectorResponse Response Object
type ShowAiOpsDetectorResponse struct {

	// **参数解释**： 全量检测项，返回所有可用检测项ID、名称，以及检测描述。 **取值范围**： 不涉及
	FullDetection *[]AiOpsDetector `json:"full_detection,omitempty"`

	// **参数解释**： 集群不可用检测项，返回所有集群不可用分类的检测项ID、名称，以及检测描述。 **取值范围**： 不涉及
	UnavailabilityDetection *[]AiOpsDetector `json:"unavailability_detection,omitempty"`
	HttpStatusCode          int              `json:"-"`
}

func (o ShowAiOpsDetectorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiOpsDetectorResponse struct{}"
	}

	return strings.Join([]string{"ShowAiOpsDetectorResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArtifactsResponse Response Object
type ListArtifactsResponse struct {

	// **参数解释**： 对话列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Artifacts *[]ArtifactsRsp `json:"artifacts,omitempty"`

	// **参数解释**： 产物个数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArtifactsResponse struct{}"
	}

	return strings.Join([]string{"ListArtifactsResponse", string(data)}, " ")
}

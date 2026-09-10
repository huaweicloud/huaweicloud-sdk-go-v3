package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateArtifactsReq struct {

	// **参数解释**： 标签列表。 **约束限制**： 产物列表不能超过10条。 **取值范围**： 不涉及 **默认取值**： 不涉及
	Filenames []string `json:"filenames"`

	StageName *StageName `json:"stage_name"`
}

func (o CreateArtifactsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArtifactsReq struct{}"
	}

	return strings.Join([]string{"CreateArtifactsReq", string(data)}, " ")
}

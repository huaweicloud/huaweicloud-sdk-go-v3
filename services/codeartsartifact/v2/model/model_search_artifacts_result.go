package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchArtifactsResult struct {

	// **参数解释**： 文件列表。 **取值范围**： 不涉及。
	Artifacts *[]ArtifactSearchResult `json:"artifacts,omitempty"`
}

func (o SearchArtifactsResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchArtifactsResult struct{}"
	}

	return strings.Join([]string{"SearchArtifactsResult", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedCommitSimpleDto struct {

	// 关联工作项ID
	RelatedId *string `json:"related_id,omitempty"`

	// 关联工作URL
	RelatedUrl *string `json:"related_url,omitempty"`
}

func (o RelatedCommitSimpleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedCommitSimpleDto struct{}"
	}

	return strings.Join([]string{"RelatedCommitSimpleDto", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedCommitListVo struct {
	Total *int32 `json:"total,omitempty"`

	List *[]RelatedCommitVo `json:"list,omitempty"`
}

func (o RelatedCommitListVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedCommitListVo struct{}"
	}

	return strings.Join([]string{"RelatedCommitListVo", string(data)}, " ")
}

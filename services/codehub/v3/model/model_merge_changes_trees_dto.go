package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MergeChangesTreesDto struct {
	Tree *[]MergeChangesTrees `json:"tree,omitempty"`
}

func (o MergeChangesTreesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeChangesTreesDto struct{}"
	}

	return strings.Join([]string{"MergeChangesTreesDto", string(data)}, " ")
}

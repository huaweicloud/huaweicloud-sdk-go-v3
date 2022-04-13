package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowVocabulariesParams struct {
	// 热词表的表名，用于筛选热词表表名。

	Name *string `json:"name,omitempty"`
}

func (o ShowVocabulariesParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVocabulariesParams struct{}"
	}

	return strings.Join([]string{"ShowVocabulariesParams", string(data)}, " ")
}

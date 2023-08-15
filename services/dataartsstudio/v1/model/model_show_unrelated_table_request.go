package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUnrelatedTableRequest Request Object
type ShowUnrelatedTableRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	Body *SearchParametersExt `json:"body,omitempty"`
}

func (o ShowUnrelatedTableRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUnrelatedTableRequest struct{}"
	}

	return strings.Join([]string{"ShowUnrelatedTableRequest", string(data)}, " ")
}

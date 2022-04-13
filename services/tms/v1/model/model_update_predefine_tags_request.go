package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePredefineTagsRequest struct {
	Body *ModifyPrefineTag `json:"body,omitempty"`
}

func (o UpdatePredefineTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"UpdatePredefineTagsRequest", string(data)}, " ")
}

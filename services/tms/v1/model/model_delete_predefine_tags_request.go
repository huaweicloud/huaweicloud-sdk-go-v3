package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeletePredefineTagsRequest struct {
	Body *ReqDeletePredefineTag `json:"body,omitempty"`
}

func (o DeletePredefineTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"DeletePredefineTagsRequest", string(data)}, " ")
}

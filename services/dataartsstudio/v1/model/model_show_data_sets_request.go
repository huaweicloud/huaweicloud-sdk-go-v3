package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDataSetsRequest Request Object
type ShowDataSetsRequest struct {

	// 实例id
	Instance string `json:"instance"`

	Body *SearchParameter `json:"body,omitempty"`
}

func (o ShowDataSetsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDataSetsRequest struct{}"
	}

	return strings.Join([]string{"ShowDataSetsRequest", string(data)}, " ")
}

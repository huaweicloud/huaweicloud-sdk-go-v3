package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddOuRequest Request Object
type AddOuRequest struct {
	Body *AddOuNameInfoV2Req `json:"body,omitempty"`
}

func (o AddOuRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOuRequest struct{}"
	}

	return strings.Join([]string{"AddOuRequest", string(data)}, " ")
}

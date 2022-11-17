package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddServiceSetUsingPostRequest struct {
	Body *AddServiceSetUsingPostRequestBody `json:"body,omitempty"`
}

func (o AddServiceSetUsingPostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddServiceSetUsingPostRequest struct{}"
	}

	return strings.Join([]string{"AddServiceSetUsingPostRequest", string(data)}, " ")
}

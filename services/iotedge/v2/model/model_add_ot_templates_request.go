package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddOtTemplatesRequest Request Object
type AddOtTemplatesRequest struct {
	Body *CreateOtTemplatesReqDto `json:"body,omitempty"`
}

func (o AddOtTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddOtTemplatesRequest struct{}"
	}

	return strings.Join([]string{"AddOtTemplatesRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdmetWithCustomPropsRequest Request Object
type ShowAdmetWithCustomPropsRequest struct {
	Body *AdmetWithCustomRequest `json:"body,omitempty"`
}

func (o ShowAdmetWithCustomPropsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdmetWithCustomPropsRequest struct{}"
	}

	return strings.Join([]string{"ShowAdmetWithCustomPropsRequest", string(data)}, " ")
}

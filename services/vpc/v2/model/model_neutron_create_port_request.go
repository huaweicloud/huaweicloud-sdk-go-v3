package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NeutronCreatePortRequest Request Object
type NeutronCreatePortRequest struct {
	Body *NeutronCreatePortRequestBody `json:"body,omitempty"`
}

func (o NeutronCreatePortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NeutronCreatePortRequest struct{}"
	}

	return strings.Join([]string{"NeutronCreatePortRequest", string(data)}, " ")
}

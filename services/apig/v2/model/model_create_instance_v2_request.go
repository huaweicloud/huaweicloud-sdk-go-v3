package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceV2Request Request Object
type CreateInstanceV2Request struct {
	Body *InstanceCreateReq `json:"body,omitempty"`
}

func (o CreateInstanceV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceV2Request struct{}"
	}

	return strings.Join([]string{"CreateInstanceV2Request", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAppV3Request struct {
	Body *CreateAppRequestBody `json:"body,omitempty"`
}

func (o CreateAppV3Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppV3Request struct{}"
	}

	return strings.Join([]string{"CreateAppV3Request", string(data)}, " ")
}

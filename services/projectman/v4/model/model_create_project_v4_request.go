package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateProjectV4Request Request Object
type CreateProjectV4Request struct {
	Body *CreateProjectV4RequestBody `json:"body,omitempty"`
}

func (o CreateProjectV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectV4Request struct{}"
	}

	return strings.Join([]string{"CreateProjectV4Request", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIterationV4Request Request Object
type CreateIterationV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	Body *CreateIterationRequestV4 `json:"body,omitempty"`
}

func (o CreateIterationV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIterationV4Request struct{}"
	}

	return strings.Join([]string{"CreateIterationV4Request", string(data)}, " ")
}

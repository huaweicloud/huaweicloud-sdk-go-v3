package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIterationV4Request Request Object
type UpdateIterationV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 迭代id
	IterationId int32 `json:"iteration_id"`

	Body *UpdateIterationRequestV4 `json:"body,omitempty"`
}

func (o UpdateIterationV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIterationV4Request struct{}"
	}

	return strings.Join([]string{"UpdateIterationV4Request", string(data)}, " ")
}

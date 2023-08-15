package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIterationV4Request Request Object
type DeleteIterationV4Request struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 迭代id
	IterationId int32 `json:"iteration_id"`
}

func (o DeleteIterationV4Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIterationV4Request struct{}"
	}

	return strings.Join([]string{"DeleteIterationV4Request", string(data)}, " ")
}

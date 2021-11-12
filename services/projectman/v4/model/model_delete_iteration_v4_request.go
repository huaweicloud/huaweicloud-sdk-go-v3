package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteIterationV4Request struct {
	// devcloud的项目id

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

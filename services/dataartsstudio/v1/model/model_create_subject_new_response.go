package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubjectNewResponse Response Object
type CreateSubjectNewResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateSubjectNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubjectNewResponse struct{}"
	}

	return strings.Join([]string{"CreateSubjectNewResponse", string(data)}, " ")
}

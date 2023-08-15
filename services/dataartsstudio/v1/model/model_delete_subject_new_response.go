package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubjectNewResponse Response Object
type DeleteSubjectNewResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteSubjectNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubjectNewResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubjectNewResponse", string(data)}, " ")
}

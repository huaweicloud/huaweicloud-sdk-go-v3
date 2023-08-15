package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportTrafficResponse Response Object
type ImportTrafficResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表
	Jobs           *[]PhoneJob `json:"jobs,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ImportTrafficResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportTrafficResponse struct{}"
	}

	return strings.Join([]string{"ImportTrafficResponse", string(data)}, " ")
}

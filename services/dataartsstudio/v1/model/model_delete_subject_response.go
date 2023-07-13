package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSubjectResponse Response Object
type DeleteSubjectResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteSubjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSubjectResponse struct{}"
	}

	return strings.Join([]string{"DeleteSubjectResponse", string(data)}, " ")
}

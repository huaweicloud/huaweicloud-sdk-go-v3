package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchSubjectResponse Response Object
type SearchSubjectResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchSubjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchSubjectResponse struct{}"
	}

	return strings.Join([]string{"SearchSubjectResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchApprovalsResponse Response Object
type SearchApprovalsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o SearchApprovalsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApprovalsResponse struct{}"
	}

	return strings.Join([]string{"SearchApprovalsResponse", string(data)}, " ")
}

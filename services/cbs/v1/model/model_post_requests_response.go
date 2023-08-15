package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PostRequestsResponse Response Object
type PostRequestsResponse struct {

	// 请求ID。  调用失败时无此字段。
	RequestId *string `json:"request_id,omitempty"`

	// 问题。  调用失败时无此字段。
	Question *string `json:"question,omitempty"`

	// 最相似的问题集。 调用失败时无此字段。
	Answers *[]Answers `json:"answers,omitempty"`

	Extends        *OldExtends `json:"extends,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o PostRequestsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostRequestsResponse struct{}"
	}

	return strings.Join([]string{"PostRequestsResponse", string(data)}, " ")
}

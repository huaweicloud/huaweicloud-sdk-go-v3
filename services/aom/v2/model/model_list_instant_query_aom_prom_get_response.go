package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstantQueryAomPromGetResponse Response Object
type ListInstantQueryAomPromGetResponse struct {

	// 响应状态。
	Status *string `json:"status,omitempty"`

	// 响应数据。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListInstantQueryAomPromGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstantQueryAomPromGetResponse struct{}"
	}

	return strings.Join([]string{"ListInstantQueryAomPromGetResponse", string(data)}, " ")
}

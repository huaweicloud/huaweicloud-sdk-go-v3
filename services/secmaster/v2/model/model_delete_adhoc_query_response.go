package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAdhocQueryResponse Response Object
type DeleteAdhocQueryResponse struct {

	// Http状态码
	Status *int32 `json:"status,omitempty"`

	// 返回信息
	Msg            *string `json:"msg,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAdhocQueryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAdhocQueryResponse struct{}"
	}

	return strings.Join([]string{"DeleteAdhocQueryResponse", string(data)}, " ")
}

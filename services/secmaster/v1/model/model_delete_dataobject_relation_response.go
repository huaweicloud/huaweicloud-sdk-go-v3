package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataobjectRelationResponse Response Object
type DeleteDataobjectRelationResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *BatchOperateDataobjectResult `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataobjectRelationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataobjectRelationResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataobjectRelationResponse", string(data)}, " ")
}

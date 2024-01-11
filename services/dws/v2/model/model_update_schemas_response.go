package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSchemasResponse Response Object
type UpdateSchemasResponse struct {

	// 响应编码。
	RetCode        *int32 `json:"ret_code,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateSchemasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSchemasResponse struct{}"
	}

	return strings.Join([]string{"UpdateSchemasResponse", string(data)}, " ")
}

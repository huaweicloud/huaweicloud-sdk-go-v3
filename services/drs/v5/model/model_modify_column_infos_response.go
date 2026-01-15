package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyColumnInfosResponse Response Object
type ModifyColumnInfosResponse struct {

	// 修改结果
	ModifyResult   *bool `json:"modify_result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ModifyColumnInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyColumnInfosResponse struct{}"
	}

	return strings.Join([]string{"ModifyColumnInfosResponse", string(data)}, " ")
}

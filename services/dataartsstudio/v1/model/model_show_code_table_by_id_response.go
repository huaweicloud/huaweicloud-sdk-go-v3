package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCodeTableByIdResponse Response Object
type ShowCodeTableByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowCodeTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCodeTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowCodeTableByIdResponse", string(data)}, " ")
}

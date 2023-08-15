package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCodeTableResponse Response Object
type DeleteCodeTableResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteCodeTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCodeTableResponse struct{}"
	}

	return strings.Join([]string{"DeleteCodeTableResponse", string(data)}, " ")
}

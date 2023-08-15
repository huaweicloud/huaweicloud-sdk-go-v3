package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCodeTableResponse Response Object
type UpdateCodeTableResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateCodeTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeTableResponse struct{}"
	}

	return strings.Join([]string{"UpdateCodeTableResponse", string(data)}, " ")
}

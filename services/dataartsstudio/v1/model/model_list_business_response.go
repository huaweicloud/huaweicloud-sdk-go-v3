package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBusinessResponse Response Object
type ListBusinessResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListBusinessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessResponse struct{}"
	}

	return strings.Join([]string{"ListBusinessResponse", string(data)}, " ")
}

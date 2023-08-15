package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTableModelResponse Response Object
type DeleteTableModelResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o DeleteTableModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTableModelResponse struct{}"
	}

	return strings.Join([]string{"DeleteTableModelResponse", string(data)}, " ")
}

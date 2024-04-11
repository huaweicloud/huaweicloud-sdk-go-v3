package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTableModelResponse Response Object
type CreateTableModelResponse struct {

	// 返回的数据信息。
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CreateTableModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTableModelResponse struct{}"
	}

	return strings.Join([]string{"CreateTableModelResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTableModelResponse Response Object
type UpdateTableModelResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateTableModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTableModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateTableModelResponse", string(data)}, " ")
}

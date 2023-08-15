package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAllModelsResponse Response Object
type CountAllModelsResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o CountAllModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAllModelsResponse struct{}"
	}

	return strings.Join([]string{"CountAllModelsResponse", string(data)}, " ")
}

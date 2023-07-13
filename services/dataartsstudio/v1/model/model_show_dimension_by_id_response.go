package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDimensionByIdResponse Response Object
type ShowDimensionByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowDimensionByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDimensionByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowDimensionByIdResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFactLogicTableByIdResponse Response Object
type ShowFactLogicTableByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowFactLogicTableByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFactLogicTableByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowFactLogicTableByIdResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SuccessRspData 响应体
type SuccessRspData struct {

	// 标识ID
	Id *string `json:"id,omitempty"`
}

func (o SuccessRspData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SuccessRspData struct{}"
	}

	return strings.Join([]string{"SuccessRspData", string(data)}, " ")
}

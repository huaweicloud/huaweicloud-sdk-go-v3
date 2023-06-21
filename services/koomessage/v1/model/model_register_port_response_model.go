package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 登记通道号返回体。
type RegisterPortResponseModel struct {
	Data *RegisterResult `json:"data,omitempty"`
}

func (o RegisterPortResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterPortResponseModel struct{}"
	}

	return strings.Join([]string{"RegisterPortResponseModel", string(data)}, " ")
}

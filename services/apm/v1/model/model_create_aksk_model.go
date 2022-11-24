package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 新增ak-sk入参。
type CreateAkskModel struct {

	// 描述信息。
	Descp *string `json:"descp,omitempty"`
}

func (o CreateAkskModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAkskModel struct{}"
	}

	return strings.Join([]string{"CreateAkskModel", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecycleBinRequestBody struct {

	// 是否启用回收站。  取值： - true：启用回收站。 - false：不启用回收站。
	Enable *bool `json:"enable,omitempty"`
}

func (o RecycleBinRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinRequestBody struct{}"
	}

	return strings.Join([]string{"RecycleBinRequestBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UploadKieRespFailure struct {

	// 配置项的key值
	Key *string `json:"key,omitempty"`

	// 配置项的labels值
	Labels *interface{} `json:"labels,omitempty"`

	// 导入失败的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 导入失败的原因
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o UploadKieRespFailure) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKieRespFailure struct{}"
	}

	return strings.Join([]string{"UploadKieRespFailure", string(data)}, " ")
}

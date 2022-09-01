package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadKieResponse struct {

	// 导入成功的配置项列表。
	Success *[]GetKieConfigs `json:"success,omitempty" xml:"success"`

	// 导入失败的配置项及其错误列表。
	Failure        *[]UploadKieRespFailure `json:"failure,omitempty" xml:"failure"`
	HttpStatusCode int                     `json:"-"`
}

func (o UploadKieResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadKieResponse struct{}"
	}

	return strings.Join([]string{"UploadKieResponse", string(data)}, " ")
}

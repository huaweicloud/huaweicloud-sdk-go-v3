package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppCodeAutoV2Response struct {

	// App Code值  支持英文，+_!@#$%+/=，且只能以英文和+、/开头，64-180个字符。
	AppCode string `json:"app_code" xml:"app_code"`

	// 编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 创建时间
	CreateTime     *string `json:"create_time,omitempty" xml:"create_time"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppCodeAutoV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppCodeAutoV2Response struct{}"
	}

	return strings.Join([]string{"CreateAppCodeAutoV2Response", string(data)}, " ")
}

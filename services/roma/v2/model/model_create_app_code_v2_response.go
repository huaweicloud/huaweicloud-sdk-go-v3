package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateAppCodeV2Response struct {

	// App Code值  支持英文，+_!@#$%+/=，且只能以英文和+、/开头。
	AppCode string `json:"app_code" xml:"app_code"`

	// 编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 创建时间
	CreateTime     *sdktime.SdkTime `json:"create_time,omitempty" xml:"create_time"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateAppCodeV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppCodeV2Response struct{}"
	}

	return strings.Join([]string{"CreateAppCodeV2Response", string(data)}, " ")
}

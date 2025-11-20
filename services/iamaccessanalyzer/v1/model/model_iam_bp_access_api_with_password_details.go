package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IamBpAccessApiWithPasswordDetails 使用密码访问API分析详细结果。
type IamBpAccessApiWithPasswordDetails struct {

	// 用户的唯一标识符（ID）。
	UserId string `json:"user_id"`

	// 用户使用密码访问API的最后时间。
	LastAccessApiWithPwdAt *sdktime.SdkTime `json:"last_access_api_with_pwd_at"`

	// 用户的创建时间。
	UserCreatedAt *sdktime.SdkTime `json:"user_created_at"`
}

func (o IamBpAccessApiWithPasswordDetails) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IamBpAccessApiWithPasswordDetails struct{}"
	}

	return strings.Join([]string{"IamBpAccessApiWithPasswordDetails", string(data)}, " ")
}

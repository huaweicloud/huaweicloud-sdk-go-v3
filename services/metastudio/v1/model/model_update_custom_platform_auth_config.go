package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCustomPlatformAuthConfig 自定义直播平台鉴权配置
type UpdateCustomPlatformAuthConfig struct {

	// 密钥Key。调用自定义直播平台时使用。 使用方式： 调用接口时，增加两个头域：x-hw-mss-time，x-hw-mss-secret * x-hw-mss-time：当前时间戳。Unix时间戳的十六进制结果。1分钟内有效。   示例： 66df9308（即2024.09.10 08:30:00） * x-hw-mss-secret：鉴权串。hmac_sha256(Key, URI（product_query_url,query参数按照Key的字典序排列）+ x-hw-mss-time)  示例： URL  https://api.example.com/v1/products?live_id=111&limit=10&offset=0   Key：GCTbw44s6MPLh4GqgDpnfuFHgy25Enly   hwTime：66df9308   x-hw-mss-secret=hmac_sha256(GCTbw44s6MPLh4GqgDpnfuFHgy25Enly,api.example.com/v1/products?limit=10&live_id=111&offset=066df9308)=5e7fe8869a12a87ea966d9edbc02e38cd4ce62c73b8b05c638f15835e78902d7   x-hw-mss-secret: 5e7fe8869a12a87ea966d9edbc02e38cd4ce62c73b8b05c638f15835e78902d7
	Key *string `json:"key,omitempty"`
}

func (o UpdateCustomPlatformAuthConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCustomPlatformAuthConfig struct{}"
	}

	return strings.Join([]string{"UpdateCustomPlatformAuthConfig", string(data)}, " ")
}

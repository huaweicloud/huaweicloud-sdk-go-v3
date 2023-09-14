package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVersionRequestBody 更新凭据对象的元数据信息请求体
type UpdateVersionRequestBody struct {

	// 凭据版本过期时间，时间戳，即从1970年1月1日至该时间的总秒数。默认为空，凭据订阅“版本过期”事件类型时，有效期判断所依据的值。
	ExpireTime int64 `json:"expire_time"`
}

func (o UpdateVersionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVersionRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVersionRequestBody", string(data)}, " ")
}

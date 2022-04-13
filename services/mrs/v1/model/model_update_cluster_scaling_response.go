package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateClusterScalingResponse struct {
	// 操作结果。 - succeeded：操作成功 - 操作失败时返回的错误码信息如[错误码](https://support.huaweicloud.com/api-mrs/ErrorCode.html)所示。

	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateClusterScalingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClusterScalingResponse struct{}"
	}

	return strings.Join([]string{"UpdateClusterScalingResponse", string(data)}, " ")
}

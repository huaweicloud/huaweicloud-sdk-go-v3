/*
 * EIP
 *
 * 云服务接口
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// 创建共享带宽请求体
type CreateSharedBandwidhRequestBody struct {
	Bandwidth *CreateSharedBandwidthOption `json:"bandwidth"`
}

func (o CreateSharedBandwidhRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateSharedBandwidhRequestBody", string(data)}, " ")
}

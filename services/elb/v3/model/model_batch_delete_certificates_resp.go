package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteCertificatesResp struct {

	// 证书ID。
	Id string `json:"id"`

	// 对应id的证书删除后的结果，not found表示证书不存在，successful表示删除成功
	RetStatus string `json:"ret_status"`

	// 错误码，删除失败时返回此字段
	RetCode *string `json:"ret_code,omitempty"`
}

func (o BatchDeleteCertificatesResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCertificatesResp struct{}"
	}

	return strings.Join([]string{"BatchDeleteCertificatesResp", string(data)}, " ")
}

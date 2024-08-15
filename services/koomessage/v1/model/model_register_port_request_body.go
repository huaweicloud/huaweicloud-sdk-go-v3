package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterPortRequestBody 通道号登记请求体。
type RegisterPortRequestBody struct {

	// 通道号。 - port_type=5时 ，长度必须为5 - port_type=1或3，长度在21位内
	Port string `json:"port"`

	// 通道号类型。 - 1：普通 - 3：前缀号段 - 5：后缀号段
	PortType int32 `json:"port_type"`

	// 签名列表，最大长度为5。单个签名长度为2-18。
	Sign []string `json:"sign"`

	// 是否需要校验。  - 0：不校验  - 1：校验签名  > 当port_type为3或者5时，sign_check必须为1。
	SignCheck int32 `json:"sign_check"`

	// 授权证明图片资源，支持jpg、bmp、png和jpeg格式，全部图片总大小不超过4M，最多支持5张。参数格式为：*资源ID:资源URL*，样例：3d214a61672846f88ad77597f935cccc:AimSauploadService/272957b708ac4891a6d5282ccd2175cccc.png。 > 资源ID与资源URL对应上传智能信息服务号图片资源API返回的resource_id和resource_url。
	AuthorizationFiles []string `json:"authorization_files"`
}

func (o RegisterPortRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterPortRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterPortRequestBody", string(data)}, " ")
}

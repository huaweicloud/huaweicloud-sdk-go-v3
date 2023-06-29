package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VendorDto vendor配置请求体
type VendorDto struct {

	// 供应商id
	Id string `json:"id"`

	// 账号id
	DomainId string `json:"domain_id"`

	// 名称
	Name string `json:"name"`

	// logo图片base64编码
	Logo string `json:"logo"`
}

func (o VendorDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VendorDto struct{}"
	}

	return strings.Join([]string{"VendorDto", string(data)}, " ")
}

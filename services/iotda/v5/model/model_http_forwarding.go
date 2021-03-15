package model

import (
	"encoding/json"

	"strings"
)

// httpserver配置信息
type HttpForwarding struct {
	// 用于接收满足规则条件数据的http服务器地址。
	Url string `json:"url"`
	// 证书id，通过证书上传接口上传证书获取Id
	CertId *string `json:"cert_id,omitempty"`
	// 当sni_enable为true时，此字段需要填写，当sni_enbale为false时，此字段默认为*
	CnName *string `json:"cn_name,omitempty"`
	// 需要https服务端和客户端都支持此功能，默认为false，设成true表明Https的客户端在发起请求时，需要携带cn_name；https服务端根据cn_name返回对应的证书；设为false可关闭此功能
	SniEnable *bool `json:"sni_enable,omitempty"`
}

func (o HttpForwarding) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HttpForwarding struct{}"
	}

	return strings.Join([]string{"HttpForwarding", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAssetTempAuthorityResponse struct {
	// 带授权签名字符串的URL。样例： ``` https://{obs_domain}/{bucket}?AWSAccessKeyId={AccessKeyID}&Expires={ExpiresValue}&Signature={Signature}

	SignStr        *string `json:"sign_str,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetTempAuthorityResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAssetTempAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetTempAuthorityResponse", string(data)}, " ")
}

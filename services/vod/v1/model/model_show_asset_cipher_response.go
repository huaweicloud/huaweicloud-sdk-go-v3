package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAssetCipherResponse struct {
	// 媒资id<br/>

	AssetId *string `json:"asset_id,omitempty"`
	// 密钥密文。未加密、正在加密或加密失败的媒资不会包含此字段。<br/>

	Edk *string `json:"edk,omitempty"`
	// 密钥明文。未加密、正在加密或加密失败的媒资不会包含此字段。<br/>

	Dk             *string `json:"dk,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetCipherResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAssetCipherResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetCipherResponse", string(data)}, " ")
}

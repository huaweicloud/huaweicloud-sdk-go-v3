package model

import (
	"encoding/json"

	"strings"
)

type CreatePreheatingAssetReq struct {
	// 根据媒资ID预热时，必选。

	AssetId *string `json:"asset_id,omitempty"`
	// 根据url预热时，必选。<br/>

	Urls *[]string `json:"urls,omitempty"`
}

func (o CreatePreheatingAssetReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePreheatingAssetReq struct{}"
	}

	return strings.Join([]string{"CreatePreheatingAssetReq", string(data)}, " ")
}

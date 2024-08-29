package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAssetReplicationInfoResponse Response Object
type ShowAssetReplicationInfoResponse struct {

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 加密后的资产信息。
	AssetInfo *string `json:"asset_info,omitempty"`

	EncryptionInfo *ReplicationEncInfo `json:"encryption_info,omitempty"`

	// 过期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetReplicationInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetReplicationInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetReplicationInfoResponse", string(data)}, " ")
}

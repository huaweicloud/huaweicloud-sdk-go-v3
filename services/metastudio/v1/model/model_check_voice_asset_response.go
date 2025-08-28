package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckVoiceAssetResponse Response Object
type CheckVoiceAssetResponse struct {

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资产名称。
	AssetName *string `json:"asset_name,omitempty"`

	// 资产类型。  公共资产类型： * VOICE_MODEL：音色模型
	AssetType *string `json:"asset_type,omitempty"`

	// 资产状态。 * CREATING：资产创建中，主文件尚未上传 * FAILED：主文件上传失败 * UNACTIVED：主文件上传成功，资产未激活，资产不可用于其他业务（用户可更新状态） * ACTIVED：主文件上传成功，资产激活，资产可用于其他业务（用户可更新状态） * DELETING：资产删除中，资产不可用，资产可恢复 * DELETED：资产文件已删除，资产不可用，资产不可恢复 * BLOCK: 资产被冻结，资产不可用，不可查看文件。
	AssetState *string `json:"asset_state,omitempty"`

	// 资产创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 资产更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	AssetExtraMeta *TtscAssetExtraMeta `json:"asset_extra_meta,omitempty"`

	// 资产下的文件。
	Files          *[]TtscAssetFileInfo `json:"files,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o CheckVoiceAssetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckVoiceAssetResponse struct{}"
	}

	return strings.Join([]string{"CheckVoiceAssetResponse", string(data)}, " ")
}

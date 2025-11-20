package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransferAssetReq 资产转移请求
type TransferAssetReq struct {
	TransferType *TransferTypeEnum `json:"transfer_type,omitempty"`

	// 资产ID列表
	AssetIds []string `json:"asset_ids"`

	// 目标用户ID
	DestProjectId string `json:"dest_project_id"`

	// 备注信息
	Memo *string `json:"memo,omitempty"`

	// 是否自动接收,管理员可用
	AutoAccept *bool `json:"auto_accept,omitempty"`

	// 是否自动激活,管理员可用
	AutoActive *bool `json:"auto_active,omitempty"`

	// 资产转移时，是否需要从接收方扣减资源（产生计费话单）
	IsNeedBilling *bool `json:"is_need_billing,omitempty"`

	// 转移任务ID，仅在transfer_type=TRANSFER_BACK时需要填写。
	TransferJobId *string `json:"transfer_job_id,omitempty"`
}

func (o TransferAssetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransferAssetReq struct{}"
	}

	return strings.Join([]string{"TransferAssetReq", string(data)}, " ")
}

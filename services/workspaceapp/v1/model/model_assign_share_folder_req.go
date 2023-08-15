package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssignShareFolderReq 共享目录分配关联对象
type AssignShareFolderReq struct {

	// WKS存储目录声明ID
	StorageClaimId string `json:"storage_claim_id"`

	// 增加的成员列表
	AddItems *[]Assignment `json:"add_items,omitempty"`

	// 移除的成员列表
	DelItems *[]Attachment `json:"del_items,omitempty"`
}

func (o AssignShareFolderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssignShareFolderReq struct{}"
	}

	return strings.Join([]string{"AssignShareFolderReq", string(data)}, " ")
}

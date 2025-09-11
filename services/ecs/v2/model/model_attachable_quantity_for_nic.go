package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AttachableQuantityForNic struct {
	FreeEfiNic *int32 `json:"free_efi_nic,omitempty"`

	FreeScsi *int32 `json:"free_scsi,omitempty"`

	FreeBlk *int32 `json:"free_blk,omitempty"`

	FreeDisk *int32 `json:"free_disk,omitempty"`

	FreeNic *int32 `json:"free_nic,omitempty"`
}

func (o AttachableQuantityForNic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachableQuantityForNic struct{}"
	}

	return strings.Join([]string{"AttachableQuantityForNic", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteCloudStorageReq 根据storage_id删除对应的文件夹。
type BatchDeleteCloudStorageReq struct {

	// cloud_storage_id,数量区间 [1, 50]。
	Items []string `json:"items"`
}

func (o BatchDeleteCloudStorageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteCloudStorageReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteCloudStorageReq", string(data)}, " ")
}

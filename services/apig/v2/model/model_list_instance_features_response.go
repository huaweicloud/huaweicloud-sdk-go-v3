package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceFeaturesResponse Response Object
type ListInstanceFeaturesResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 实例支持的特性列表： - \"resize_huge_flavor\" - \"health_check_in_instance_etcd\" - \"shubao_support_add_node\" - \"upgrade_uninterrupted\" - \"sm_cipher_type\"  与实例版本有关，列表中不展示的特性为实例不支持的特性
	Features       *[]string `json:"features,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListInstanceFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceFeaturesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceFeaturesResponse", string(data)}, " ")
}

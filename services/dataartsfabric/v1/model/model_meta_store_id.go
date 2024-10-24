package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetaStoreId Metastore信息，LakeFormation服务的实例Id，即MetaStoreId。
type MetaStoreId struct {
}

func (o MetaStoreId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaStoreId struct{}"
	}

	return strings.Join([]string{"MetaStoreId", string(data)}, " ")
}

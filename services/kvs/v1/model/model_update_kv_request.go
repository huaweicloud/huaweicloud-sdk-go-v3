package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKvRequest Request Object
type UpdateKvRequest struct {

	// 仓名，全域唯一，不同租户的仓名不能相同。  - 格式：${prefix}-${region-code}-${account-id}，其中prefix为自定义前缀，region-code为kvs集群所在的区域代码，account-id为用户的账户id - 取值字符限制：[a-z0-9-]+ - 长度：[16,52] > \"-\"不能出现在名字头部或尾部
	StoreName *string `bson:"store_name,omitempty"`

	Body *UpdateKvRequestBody `bson:"body,omitempty"`
}

func (o UpdateKvRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKvRequest struct{}"
	}

	return strings.Join([]string{"UpdateKvRequest", string(data)}, " ")
}

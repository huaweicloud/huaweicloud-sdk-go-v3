package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFreeResourceUsagesReq struct {

	// 资源项ID列表，每个最大64字节。 资源项ID，一个资源包中会含有多个资源项，一个使用量类型对应一个资源项。资源项ID来自查询资源包列表接口的响应。
	FreeResourceIds []string `json:"free_resource_ids"`
}

func (o ListFreeResourceUsagesReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourceUsagesReq struct{}"
	}

	return strings.Join([]string{"ListFreeResourceUsagesReq", string(data)}, " ")
}

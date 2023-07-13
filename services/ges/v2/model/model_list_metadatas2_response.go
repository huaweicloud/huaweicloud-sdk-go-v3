package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetadatas2Response Response Object
type ListMetadatas2Response struct {

	// 元数据返回个数。请求失败时，字段为空。
	SchemaCount *int32 `json:"schema_count,omitempty"`

	// 当前projectId下的所有元数据列表。请求失败时，字段为空。
	SchemaList     *[]ListMetadatasRespSchemaList `json:"schema_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListMetadatas2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetadatas2Response struct{}"
	}

	return strings.Join([]string{"ListMetadatas2Response", string(data)}, " ")
}

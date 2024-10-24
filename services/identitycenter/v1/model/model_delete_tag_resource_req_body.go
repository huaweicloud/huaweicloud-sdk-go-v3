package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagResourceReqBody DeleteTagResource 操作的请求体。
type DeleteTagResourceReqBody struct {

	// 用于管理资源的一组键值对
	Tags []DeleteTagDto `json:"tags"`
}

func (o DeleteTagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagResourceReqBody struct{}"
	}

	return strings.Join([]string{"DeleteTagResourceReqBody", string(data)}, " ")
}

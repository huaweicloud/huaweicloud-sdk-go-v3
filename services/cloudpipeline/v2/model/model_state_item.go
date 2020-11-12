/*
 * CloudPipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 编排State详情，map类型数据。本字段为描述流水线基础编排数据之一，建议可通过流水线真实界面基于模板创建接口中获取
type StateItem struct {
}

func (o StateItem) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"StateItem", string(data)}, " ")
}

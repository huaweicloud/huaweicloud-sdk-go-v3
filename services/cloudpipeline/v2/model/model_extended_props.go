/*
 * devcloudpipeline
 *
 * devcloud pipeline api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 任务参数,,map类型数据
type ExtendedProps struct {
}

func (o ExtendedProps) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ExtendedProps", string(data)}, " ")
}

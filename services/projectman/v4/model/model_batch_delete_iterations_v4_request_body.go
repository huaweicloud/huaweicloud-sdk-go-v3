/*
 * ProjectMan
 *
 * devcloud projectman api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type BatchDeleteIterationsV4RequestBody struct {
	// 迭代的id
	IterationIds []int32 `json:"iteration_ids"`
}

func (o BatchDeleteIterationsV4RequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"BatchDeleteIterationsV4RequestBody", string(data)}, " ")
}

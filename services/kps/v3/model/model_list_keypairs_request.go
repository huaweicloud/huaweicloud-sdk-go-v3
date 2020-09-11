/*
 * kps
 *
 * kps v3 版本API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListKeypairsRequest struct {
}

func (o ListKeypairsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListKeypairsRequest", string(data)}, " ")
}

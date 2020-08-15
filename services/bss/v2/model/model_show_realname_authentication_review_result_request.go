/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"
	"strings"
)

// Request Object
type ShowRealnameAuthenticationReviewResultRequest struct {
	CustomerId string `json:"customer_id"`
}

func (o ShowRealnameAuthenticationReviewResultRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowRealnameAuthenticationReviewResultRequest", string(data)}, " ")
}

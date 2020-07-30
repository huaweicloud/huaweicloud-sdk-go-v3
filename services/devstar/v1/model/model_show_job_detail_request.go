/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

// Request Object
type ShowJobDetailRequest struct {
	XLanguage string `json:"X-Language,omitempty"`
	JobId string `json:"job_id"`
}

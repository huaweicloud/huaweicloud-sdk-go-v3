/*
    * Devstar
    *
    * Devstar API
    *
*/

package model

// Request Object
type ShowTemplateDetailRequest struct {
	XLanguage string `json:"X-Language,omitempty"`
	TemplateId string `json:"template_id"`
}

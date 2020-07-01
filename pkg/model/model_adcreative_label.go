/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 标签
type AdcreativeLabel struct {
	CoordinateX    int64          `json:"coordinate_x,omitempty"`
	CoordinateY    int64          `json:"coordinate_y,omitempty"`
	Direction      LableDirection `json:"direction,omitempty"`
	Content        string         `json:"content,omitempty"`
	LandingPageUrl string         `json:"landing_page_url,omitempty"`
}
/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

//
type HistoryCpcSpec struct {
	Date      string `json:"date,omitempty"`
	ActualCpc int64  `json:"actual_cpc,omitempty"`
	TargetCpc int64  `json:"target_cpc,omitempty"`
}
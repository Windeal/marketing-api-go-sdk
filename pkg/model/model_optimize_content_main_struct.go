/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 诊断详情-优化操作明细
type OptimizeContentMainStruct struct {
	Status           OptimizeSubStatusStruct  `json:"status,omitempty"`
	OptimizeContents OptimizeSubContentStruct `json:"optimize_contents,omitempty"`
}
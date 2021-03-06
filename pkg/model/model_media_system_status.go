/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// MediaSystemStatus : 转码状态
type MediaSystemStatus string

// List of MediaSystemStatus
const (
	MediaSystemStatus_VALID   MediaSystemStatus = "MEDIA_STATUS_VALID"
	MediaSystemStatus_PENDING MediaSystemStatus = "MEDIA_STATUS_PENDING"
	MediaSystemStatus_ERROR_  MediaSystemStatus = "MEDIA_STATUS_ERROR"
)

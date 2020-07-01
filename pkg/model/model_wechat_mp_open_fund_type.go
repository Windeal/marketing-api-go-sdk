/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// WechatMpOpenFundType : 资金账户类型，当trade_type为AGENCY_CREDIT_REPAY时，此字段不返回
type WechatMpOpenFundType string

// List of WechatMpOpenFundType
const (
	WechatMpOpenFundType_UNSUPPORTED       WechatMpOpenFundType = "FUND_TYPE_UNSUPPORTED"
	WechatMpOpenFundType_CASH              WechatMpOpenFundType = "FUND_TYPE_CASH"
	WechatMpOpenFundType_GIFT              WechatMpOpenFundType = "FUND_TYPE_GIFT"
	WechatMpOpenFundType_CREDIT            WechatMpOpenFundType = "FUND_TYPE_CREDIT"
	WechatMpOpenFundType_SPECIAL_PROMOTION WechatMpOpenFundType = "FUND_TYPE_SPECIAL_PROMOTION"
	WechatMpOpenFundType_PAYMENT_DUE       WechatMpOpenFundType = "FUND_TYPE_PAYMENT_DUE"
	WechatMpOpenFundType_UNSUPPPORTED      WechatMpOpenFundType = "FUND_TYPE_UNSUPPPORTED"
	WechatMpOpenFundType_CREDIT_ROLL       WechatMpOpenFundType = "FUND_TYPE_CREDIT_ROLL"
	WechatMpOpenFundType_CREDIT_TEMPORARY  WechatMpOpenFundType = "FUND_TYPE_CREDIT_TEMPORARY"
	WechatMpOpenFundType_MINIPROGRAM       WechatMpOpenFundType = "FUND_TYPE_MINIPROGRAM"
)
package api

import "strings"

var reflect map[Code]string

// Code .
type Code string

func (c Code) Error() string {
	msg, ok := reflect[c]
	if !ok {
		return reflect[CheckErrInternalErr]
	}

	return msg
}

// ToString .
func (c Code) ToString() string {
	return string(c)
}

const (
	// OriginSuccess 兼容之前版本
	OriginSuccess Code = "0"
	Success       Code = "SP000000"
	Failed        Code = "SP010000"

	CheckErrParamsInvalid            Code = "SP020100"
	CheckErrTemplateInvalid          Code = "SP020200"
	CheckErrTemplateLack             Code = "SP020201"
	CheckErrTemplateCanNotParse      Code = "SP020202"
	CheckErrTemplateCanNotFind       Code = "SP020203"
	CheckErrBeforeFailed             Code = "SP020300"
	CheckErrBeforeFailedUnderwriting Code = "SP020301"
	CheckErrBeforeFailedApplyPeriod  Code = "SP020302"
	CheckErrAfterFailed              Code = "SP020400"
	CheckErrAfterFailedApplyFst      Code = "SP020401"
	CheckErrDuplicateRequest         Code = "SP020500"
	CheckErrInternalErr              Code = "SP020600"

	NetErrRequestFailed  Code = "SP030000"
	NetErrRequestInvalid Code = "SP030200"
	NetErrOverTime       Code = "SP030100"

	RespErrIsNull             Code = "SP050100"
	RespErrContentCanNotParse Code = "SP050200"
	RespErrHttpStatusInvalid  Code = "SP050300"
)

func init() {
	reflect = map[Code]string{
		Success: "成功",
		Failed:  "失败",

		CheckErrParamsInvalid:            "参数错误",
		CheckErrTemplateInvalid:          "模板错误",
		CheckErrTemplateLack:             "模板缺失",
		CheckErrTemplateCanNotParse:      "模板格式错误或缺失必要参数",
		CheckErrTemplateCanNotFind:       "模板获取失败",
		CheckErrBeforeFailed:             "前置条件失败",
		CheckErrBeforeFailedUnderwriting: "前置核保失败",
		CheckErrBeforeFailedApplyPeriod:  "前置申请续期失败",
		CheckErrAfterFailed:              "后置条件失败",
		CheckErrAfterFailedApplyFst:      "后置自动期缴第一期失败",
		CheckErrDuplicateRequest:         "重复请求",
		CheckErrInternalErr:              "内部异常",

		NetErrRequestFailed:  "网络请求错误",
		NetErrRequestInvalid: "无效的网络请求",
		NetErrOverTime:       "请求超时",

		RespErrIsNull:             "返回结果为空",
		RespErrContentCanNotParse: "返回内容解析失败",
		RespErrHttpStatusInvalid:  "返回http状态码异常",
	}
}

// IsSuccess 判断保司返回结果是否为成功，新code 为 Success，原code 为OriginSuccess
// 现做兼容处理，之后可在数据库中将映射状态设为启用，即可将原状态码全部替换为新状态码
func IsSuccess(code string) bool {
	return code == Success.ToString() ||
		code == OriginSuccess.ToString()
}

// IsNetErr .
func IsNetErr(code string) bool {
	return strings.HasPrefix(code, "SP03")
}

// IsRespErr .
func IsRespErr(code string) bool {
	return strings.HasPrefix(code, "SP05")
}

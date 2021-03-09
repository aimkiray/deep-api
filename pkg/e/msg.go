package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "Invalid Params",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token Auth Error",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token Timeout",
	ERROR_AUTH_TOKEN : "Token Generate Error",
	ERROR_AUTH : "Token Error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
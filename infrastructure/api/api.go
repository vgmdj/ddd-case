package api

// Reseponse the api response object
type Reseponse struct {
	RetCode    Code        `json:"ret_code"`
	RetMsg     string      `json:"ret_msg"`
	RetContent interface{} `json:"ret_content"`
}

// NewAPIReponse return the response object
func NewAPIReponse(code Code, msg string, data ...interface{}) Reseponse {
	if msg == "" {
		msg = code.Error()
	}

	var content interface{}
	if len(data) != 0 {
		content = data[0]
	}

	return Reseponse{
		RetCode:    code,
		RetMsg:     msg,
		RetContent: content,
	}
}

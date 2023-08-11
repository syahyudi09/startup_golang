package helper

type response struct{
	meta meta
	data interface{}
}

type meta struct {
	message string
	code int
	status string
}

func ApiResponse(message string, code int, status string, data interface{}) response{
	meta := meta{
		message : message,
		code: code,
		status: status,
	}	

	response := response{
		meta: meta,
		data: data,
	}
	
	return response
}


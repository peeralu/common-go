package exception

func NewE300Exception(message string) error {
	return AppError{
		Code:           "300",
		StatusCode:     422,
		ErrorMessageTh: "ข้อมูลไม่ถูกต้อง",
		ErrorMessageEn: "Incorrect data",
	}
}

func NewE301Exception(parameter string) error {
	return AppError{
		Code:           "301",
		StatusCode:     422,
		ErrorMessageTh: "ต้องการข้อมูล: " + parameter,
		ErrorMessageEn: "Required parameters: " + parameter,
	}
}

func NewE302Exception(parameter string) error {
	return AppError{
		Code:           "302",
		StatusCode:     422,
		ErrorMessageTh: "ประเภทข้อมูลผิด: " + parameter,
		ErrorMessageEn: "Incorrect data type: " + parameter,
	}
}

func NewE303Exception(parameter string) error {
	return AppError{
		Code:           "303",
		StatusCode:     422,
		ErrorMessageTh: "ข้อมูลมีขนาดเกินกำหนด: " + parameter,
		ErrorMessageEn: "Incorrect data length: " + parameter,
	}
}

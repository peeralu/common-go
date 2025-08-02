package exception

func NewE998Exception() error {
	return AppError{
		Code:           "998",
		StatusCode:     500,
		ErrorMessageTh: "ระบบปิดปรับปรุง",
		ErrorMessageEn: "System under maintenance",
	}
}

func NewE999Exception() error {
	return AppError{
		Code:           "999",
		StatusCode:     500,
		ErrorMessageTh: "ระบบขัดข้อง",
		ErrorMessageEn: "Internal server error",
	}
}

func NewUnknownException() AppError {
	return AppError{
		Code:           "999",
		StatusCode:     500,
		ErrorMessageTh: "ระบบขัดข้อง",
		ErrorMessageEn: "Internal server error",
	}
}

package exception

func NewE100Exception() error {
	return AppError{
		Code:           "100",
		StatusCode:     401,
		ErrorMessageTh: "ข้อมูล token ไม่ถูกต้อง",
		ErrorMessageEn: "Unauthorized token",
	}
}

func NewE101Exception() error {
	return AppError{
		Code:           "101",
		StatusCode:     403,
		ErrorMessageTh: "การยืนยันตัวตนผิดพลาด",
		ErrorMessageEn: "Authentication credentials were not provided",
	}
}

func NewE102Exception() error {
	return AppError{
		Code:           "102",
		StatusCode:     403,
		ErrorMessageTh: "ไม่มีสิทธิ์เรียกใช้บริการ",
		ErrorMessageEn: "No permission to call this service",
	}
}

func NewE103Exception() error {
	return AppError{
		Code:           "103",
		StatusCode:     404,
		ErrorMessageTh: "ไม่พบข้อมูล",
		ErrorMessageEn: "Data not found",
	}
}

func NewE104Exception() error {
	return AppError{
		Code:           "104",
		StatusCode:     405,
		ErrorMessageTh: "ไม่สามารถดำเนินการได้",
		ErrorMessageEn: "Method not allowed",
	}
}

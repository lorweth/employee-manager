package model

// Lập trình viên: mã nhân viên, họ tên, tuổi, số điện thoại, email, lương cơ bản, số giờ overtime.
// Kiểm chứng viên: mã nhân viên, họ tên, tuổi, số điện thoại, email, lương cơ bản, số lỗi phát hiện được.

type Employee interface {
	Info() AbstractEmployee
	Salary() float64
	String() string
}

type AbstractEmployee struct {
	ID       int64
	Fullname string
	Age      uint8
	PhoneNo  string
	Email    string
	BasePay  float64
}

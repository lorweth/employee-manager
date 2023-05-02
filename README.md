# Simple code for learning Generic in Go

## Requirement

Công ty phần mềm có 2 loại nhân viên: Lập trình viên là những người sẽ viết mã nguồn cho các ứng dụng, Kiểm chứng viên có nhiệm vụ kiểm tra mã nguồn và chương trình mà lập trình viên viết ra để tìm các lỗi trước khi giao sản phẩm cho khách hàng.

Hiện tại, công ty lưu trữ thông tin của các loại nhân viên như sau:

Lập trình viên: mã nhân viên, họ tên, tuổi, số điện thoại, email, lương cơ bản, số giờ overtime.
Kiểm chứng viên: mã nhân viên, họ tên, tuổi, số điện thoại, email, lương cơ bản, số lỗi phát hiện được.

Do tính chất công việc khác nhau nên lương cơ bản của lập trình viên và kiểm chứng viên cũng khác nhau. Cụ thể:

```Txt
Lương (lập trình viên) = lương cơ bản + số giờ làm thêm * 200.000.
Lương (kiểm chứng viên) = lương cơ bản + số lỗi * 50.000.
```

Hãy viết chương trình thực hiện các yêu cầu sau:

1. Nhập vào danh sách nhân viên (lưu trữ trong một mảng duy nhất).
2. Liệt kê danh sách nhân viên có lương thấp hơn mức lương trung bình của các nhân viên trong công ty.

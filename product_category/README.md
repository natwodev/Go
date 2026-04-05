# Practice: Product & Category API (Golang + Gin)

Thực hành xây dựng API cơ bản với Gin Framework cho đối tượng Product và Category.

## Thời gian thực hiện
- Dự kiến: 30/03/2026 - 05/04/2026
- Thực tế: Hoàn thành 16/04/2026

## Chức năng
- CRUD Category (ID, Name)
- CRUD Product (ID, Name, Price, CategoryID)
- Validation dữ liệu bằng `binding:"required"`
- Kiểm tra CategoryID tồn tại khi tạo Product mới.

## API Endpoints

### Category
- `GET /api/v1/categories`: Lấy danh sách category
- `GET /api/v1/categories/:id`: Lấy category theo ID
- `POST /api/v1/categories`: Tạo mới category
  - Body: `{"name": "Category Name"}`
- `PUT /api/v1/categories/:id`: Cập nhật category
- `DELETE /api/v1/categories/:id`: Xóa category

### Product
- `GET /api/v1/products`: Lấy danh sách product
- `GET /api/v1/products/:id`: Lấy product theo ID
- `POST /api/v1/products`: Tạo mới product
  - Body: `{"name": "Product Name", "price": 100.5, "category_id": 1}`
- `PUT /api/v1/products/:id`: Cập nhật product
- `DELETE /api/v1/products/:id`: Xóa product

## Cách chạy
1. Cài đặt dependencies: `go mod tidy`
2. Chạy ứng dụng: `go run main.go`
3. API sẽ lắng nghe tại: `http://localhost:8080`

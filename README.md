# Go Backend Project

Backend API được xây dựng bằng Go theo kiến trúc tách lớp `controller -> service -> repository`, sử dụng PostgreSQL và GORM để quản lý dữ liệu.

## Mục tiêu dự án

- Xây dựng REST API cho các module `Category`, `Product`, `Inventory`
- Áp dụng cấu trúc dự án rõ ràng, dễ mở rộng
- Kết nối PostgreSQL và tự động migration schema khi khởi động
- Chuẩn hóa response JSON để dễ tích hợp frontend/mobile

## Công nghệ sử dụng

- Go
- PostgreSQL
- GORM
- net/http (HTTP server/router)

## Kiến trúc thư mục chính

- `go-backend-pro/cmd/api`: điểm khởi động ứng dụng
- `go-backend-pro/internal/app/controller`: xử lý HTTP request/response
- `go-backend-pro/internal/app/service`: business logic
- `go-backend-pro/internal/app/repository`: thao tác dữ liệu
- `go-backend-pro/internal/app/model`: entity và response model
- `go-backend-pro/internal/pkg/postgres`: cấu hình và kết nối database

## Tính năng hiện có

- CRUD `Category`
- CRUD `Product`
- CRUD `Inventory`
- Inventory liên kết Product bằng foreign key
- Chặn tạo inventory nếu `product_id` không tồn tại
- Cập nhật số lượng tồn kho theo nghiệp vụ:
  - `increase` (tăng số lượng)
  - `decrease` (giảm số lượng, có kiểm tra không âm kho)
- Health check endpoint

## API chính

- `GET /health`
- `GET /api/v1/ping`

- `GET /api/v1/categories`
- `POST /api/v1/categories`
- `GET /api/v1/categories/{id}`
- `PUT /api/v1/categories/{id}`
- `DELETE /api/v1/categories/{id}`

- `GET /api/v1/products`
- `POST /api/v1/products`
- `GET /api/v1/products/{id}`
- `PUT /api/v1/products/{id}`
- `DELETE /api/v1/products/{id}`

- `GET /api/v1/inventories`
- `POST /api/v1/inventories`
- `GET /api/v1/inventories/{id}`
- `DELETE /api/v1/inventories/{id}`
- `POST /api/v1/inventories/{id}/increase`
- `POST /api/v1/inventories/{id}/decrease`

## Cách chạy local

1. Cấu hình biến môi trường:

```bash
cp go-backend-pro/.env.example go-backend-pro/.env
```

2. Chạy API:

```bash
cd go-backend-pro
go run ./cmd/api
```

## Định hướng tiếp theo

- Bổ sung JWT authentication và authorization
- Thêm middleware logging/recovery chuẩn hóa
- Viết test cho service và controller
- Tách migration SQL riêng thay vì chỉ dùng automigrate

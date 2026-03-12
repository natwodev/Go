# Thiết kế Cơ sở Dữ liệu

## Các bảng chính

### 1. Bảng `users` (Người dùng)
- `id`: Serial (Primary Key)
- `username`: Varchar(50) (Unique, Not Null)
- `password`: Varchar(255) (Hash, Not Null)
- `email`: Varchar(100) (Unique, Not Null)
- `created_at`: Timestamp (Default: Now)

### 2. Bảng `posts` (Bài viết) - Ví dụ cho quan hệ
- `id`: Serial (Primary Key)
- `title`: Varchar(255) (Not Null)
- `content`: Text
- `user_id`: Integer (Foreign Key -> users.id)
- `created_at`: Timestamp

## Sơ đồ quan hệ (ERD)
- Một User có thể có nhiều Posts (1-n).
- Một Post thuộc về một User duy nhất.

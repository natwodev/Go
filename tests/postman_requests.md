# API Testing with Postman

## Endpoints

### 1. Create User
- **Method**: `POST`
- **URL**: `http://localhost:8080/users`
- **Body** (JSON):
```json
{
    "name": "Jane Doe",
    "email": "jane@example.com"
}
```

### 2. Get User
- **Method**: `GET`
- **URL**: `http://localhost:8080/users/1`

### 3. Update User
- **Method**: `PUT`
- **URL**: `http://localhost:8080/users/1`
- **Body** (JSON):
```json
{
    "name": "Jane Wilson",
    "email": "jane.wilson@example.com"
}
```

### 4. Delete User
- **Method**: `DELETE`
- **URL**: `http://localhost:8080/users/1`

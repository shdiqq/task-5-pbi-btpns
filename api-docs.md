# API Documentation

You can find detailed API documentation [here](/api-docs).

## Authentication

To access most of the API endpoints, you'll need to authenticate using JSON Web Tokens (JWT). Here's how you can obtain a JWT:

**Endpoint:** `/users/login`

**Method:** `POST`

**Request Body:**

```json
{
  "username": string,
  "password": string
}
```

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "token": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```

Include the JWT token in the Authorization header for subsequent requests:
Authorization: Bearer your_jwt_token

## Another Endpoint

**Endpoint:** `/users/register`

**Method:** `POST`

**Request Body:**

```json
{
  "username": string,
  "email": string,
  "password": string,
  "password_confirm": string
}
```

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": { 
    "id": number,
    "username": string,
    "email": string,
    "password": string,
    "createdAt": string,
    "updatedAt": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/users/logout`

**Method:** `GET`

**Response Success:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/users`

**Method:** `GET`

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": [
    {
      "id": number,
      "username": string,
      "email": string,
      "password": string,
      "createdAt": string,
      "updatedAt": string
    },
    ...
  ]
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/users/{userId}`

**Method:** `PUT`

**Request Body:**

```json
{
  "username?": string,
  "email?": string,
  "password?": string,
}
```

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "id": number,
    "username": string,
    "email": string,
    "password": string,
    "createdAt": string,
    "updatedAt": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/users/{userId}`

**Method:** `DELETE`

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "id": number,
    "username": string,
    "email": string,
    "password": string,
    "createdAt": string,
    "updatedAt": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/photos`

**Method:** `GET`

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": [
    {
      "id": number,
      "title": string,
      "caption": string,
      "photoUrl": string,
      "userId": string,
      "created_at": string,
      "updated_at": string
    },
    ...
  ]
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/photos`

**Method:** `POST`

**Request Body:**

```json
{
  "title": string,
  "caption": string,
  "photoUrl": string,
}
```

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "id": number,
    "title": string,
    "caption": string,
    "photoUrl": string,
    "userId": string,
    "created_at": string,
    "updated_at": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/photos/{photoId}`

**Method:** `GET`

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "id": number,
    "title": string,
    "caption": string,
    "photoUrl": string,
    "userId": string,
    "User": {
      "id": number,
      "username": string,
      "email": string,
      "password": string,
      "createdAt": string,
      "updatedAt": string
    },
    "created_at": string,
    "updated_at": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/photos/{photoId}`

**Method:** `PUT`

**Request Body:**

```json
{
  "title?": string,
  "caption?": string,
  "photoUrl?": string,
}
```

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "id": number,
    "title": string,
    "caption": string,
    "photoUrl": string,
    "userId": string,
    "created_at": string,
    "updated_at": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
**Endpoint:** `/photos/{photoId}`

**Method:** `DELETE`

**Response Success:**
```json
{
  "status": string,
  "message": string,
  "data": {
    "id": number,
    "title": string,
    "caption": string,
    "photoUrl": string,
    "userId": string,
    "created_at": string,
    "updated_at": string
  }
}
```

**Response Fail:**
```json
{
  "status": string,
  "message": string
}
```
---
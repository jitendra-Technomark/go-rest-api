# Wishlist API Documentation

## 1. Database Schema
The database schema for the wishlist functionality consists of a single table:

gssTBL, psTBL

### Wishlist - make sepearte for gss & ps - database/schema/m<>_yyyymmdd.sql
- wishlist table (gss):
id (primary key) int not null
user_Id int (foreign key referencing userTBL table) int not null
space_Id (foreign key referencing space table) int not null
start_Date datetime not null
end_Date datetime not null
is_add bit(1)
created_On datetime default CURRENT_TIMESTAMP
updated_On datetime CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP

- wishlist table (ps):
id (primary key) int not null
user_Id int (foreign key referencing users table) int not null
space_Id (foreign key referencing space table) int not null
start_Date datetime not null
end_Date datetime not null
is_add bit(1)
created_On datetime default CURRENT_TIMESTAMP
updated_On datetime CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP

***
add package wishlist
## 2. API Endpoints:
1. Add to wishlist
**Endpoint:** `/api/v1/wishlist/gss/:id` `/api/v1/wishlist/ps/:id`
**Method:** POST
**Request Body**: json
{
    "start_Date": "YYYY-MM-DD",
    "end_Date": "YYYY-MM-DD",
}

**Response**: 
- Status Code: `201 Created`
{
    "message": "Space added to wishlist"
}
- Status Code: `400 Bad Request`
{
    "message": "Invalid input data. Please check the request payload."
}
- Status Code: `404 Not Found`
{
    "message": "The requested resource was not found."
}
- Status Code: `401 Unauthorized`
{
    "message": "Authentication credentials are missing or invalid."
}

***

2. Remove Item from Wishlist
**Endpoint:** `/api/v1/wishlist/gss/:id` `/api/v1/wishlist/ps/:id`
**Method:** DELETE
**Request Body**: NA

**Response**: 
- Status Code: `200 OK`
{
    "message": "Space successfully removed from wishlist"
}
- Status Code: `400 Bad Request`
{
    "message": "Invalid input data. Please check the request payload."
}
- Status Code: `404 Not Found`
{
    "message": "The requested resource was not found."
}
- Status Code: `401 Unauthorized`
{
    "message": "Authentication credentials are missing or invalid."
}

***

3. Get Wishlist for a User ( merge list from seperate tables)
**Endpoint:** `/api/v1/wishlist`
**Method:** GET
**Request Body**: NA
**Response**: 
- Status Code: `200 OK`
[
    {
        "id": 1,
        "space_type": "string",
        "space_id": 123,
        "start_date": "YYYY-MM-DD",
        "end_date": "YYYY-MM-DD",
        "created_on": "YYYY-MM-DD"
        "updated_on": "YYYY-MM-DD",
    },
    {
        "id": 2,
        "space_type": "string",
        "space_id": 456,
        "start_date": "YYYY-MM-DD",
        "end_date": "YYYY-MM-DD",
        "created_on": "YYYY-MM-DD"
        "updated_on": "YYYY-MM-DD",
    }
]
 - Status Code: `400 Bad Request`
{
    "message": "Invalid input data. Please check the request payload."
}
- Status Code: `404 Not Found`
{
    "message": "The requested resource was not found."
}
- Status Code: `401 Unauthorized`
{
    "message": "Authentication credentials are missing or invalid."
}
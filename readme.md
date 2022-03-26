
# Notes APP

A REST API backend application that can be used to manage personal notes in a multi-user environment.


### How to start 

- Make sure database is up and running and you have update the ```configurations/database.yaml``` file with relevant values.
- Run the ```database.sql``` query to get database and table created.
- Clone the service locally and run the service by typing ```go run main.go```
- make sure service is up and running. 
- Now you can send request to the service.

Note: Use the postman collection ```notes_app.postman_collection.json```

## API Reference

#### Get all un archived notes

```http
  GET /note/get_un_archived
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `paginator` | `{"page":1, "size": 10}` | **required**. page must be greater than zero & size must be in between 10 - 100 |

- Response 
```json
    {
        "data": [
            {
                "id": 3,
                "user_id": 1,
                "title": "abc",
                "description": "cc bnbj v",
                "created_at": "2022-03-26T22:10:40Z",
                "updated_at": "2022-03-26T22:31:01Z"
            }
        ],
        "error": null,
        "success": "true"
    }
```

#### Get all archived notes

```http
  GET /note/get_archived
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `paginator` | `{"page":1, "size": 10}` | **required**. page must be greater than zero & size must be in between 10 - 100 |

- Response 
```json
    {
        "data": [
            {
                "id": 2,
                "user_id": 1,
                "title": "hello",
                "description": "cc bnbj v",
                "created_at": "2022-03-26T22:10:40Z",
                "updated_at": "2022-03-26T22:31:01Z"
            }
        ],
        "error": null,
        "success": "true"
    }
```

#### Save a note

```http
  POST /note/create
```

 - Request body


```json
    {
        "user_id": 123,
        "title": "title",
        "description": "description"
    }
```

- Response 
```json
    {
        "data": 3,
        "error": null,
        "success": "true"
    }
```

#### Archive a note

```http
  GET /note/archive/{id}
```

- Response 
```json
    {
        "data": 3,
        "error": null,
        "success": true
    }
```

#### Un Archive a note

```http
  GET /note/un_archive/{id}
```

- Response 
```json
    {
        "data": 3,
        "error": null,
        "success": true
    }
```

#### Update a note

```http
  PUT /note/update/{id}
```

 - Request body


```json
    {
        "title": "title",
        "description": "description"
    }
```

- Response 
```json
    {
        "data": 2,
        "error": null,
        "success": true
    }
```

#### Delete a note

```http
  DELETE /note/delete/{id}
```

- Response 
```json
    {
        "data": 3,
        "error": null,
        "success": true
    }
```

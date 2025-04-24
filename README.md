# Students API

Students API is a simple API written in Golang that allows you to manage students. It uses SQLite as a database.

## Installation

You can download it from the official website~
Official website: https://go.dev/doc/install

Once you have Go installed, you can clone the repository and build the application:

```bash
git clone https://github.com/techatikin/students-api.git
cd students-api
go build
```

## Configuration

The application uses a configuration file to store its settings. You can find the configuration file in the `config` directory.

Here's an example of the configuration file:

```yaml
env: "dev"
storage_path: <path_to_sqlite_db_file>
http_server:
  address: "localhost:8082"
```

You need to replace `<path_to_sqlite_db_file>` with the path where you want to store the SQLite database file.

After that, you can run the application using the following command:

```bash
go run cmd/students-api/main.go -config config/local.yaml
```

The application will start listening on port 8082. You can access the API using a tool like Postman or cURL.

## Endpoints

### Create Student

To create a new student, you need to send a POST request to the `/api/students` endpoint with the following JSON body:

```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 25
}
```

If the request is successful, the API will return a 201 Created status code and the newly created student's ID in the response body.

### Get Student By ID

To get a student by its ID, you need to send a GET request to the `/api/students/{id}` endpoint.

If the request is successful, the API will return a 200 OK status code and the student's data in the response body.

### Get All Students

To get all students, you need to send a GET request to the `/api/students` endpoint.

If the request is successful, the API will return a 200 OK status code and an array of students in the response body.

### Update Student By ID

To update a student by its ID, you need to send a PUT request to the `/api/students/{id}` endpoint with the following JSON body:

```json
{
  "name": "Jane Doe",
  "email": "jane.doe@example.com",
  "age": 30
}
```

If the request is successful, the API will return a 200 OK status code and the updated student's data in the response body.

### Delete Student By ID

To delete a student by its ID, you need to send a DELETE request to the `/api/students/{id}` endpoint.

If the request is successful, the API will return a 204 No Content status code.

## Conclusion

Students API is a simple API with CRUD operations for managing students. It uses SQLite as a database and provides a RESTful API for interacting with the database.

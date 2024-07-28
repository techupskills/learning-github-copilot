techupskills: Provide doc with examples on how to use the API in #file:EmployeeController.java 

GitHub Copilot: # Employee Directory API Documentation

The Employee Directory API provides various endpoints to search and filter employees from a directory based on different parameters such as employee ID, age, and city. Below are the available endpoints with examples on how to use them.

## Base URL

```
http://<host>:<port>/employees
```

Replace `<host>` and `<port>` with the actual host and port where the service is running.

## Endpoints

### 1. Get All Employees

- **Endpoint:** `/all`
- **Method:** GET
- **Description:** Retrieves a list of all employees.
- **Example Request:**

  ```
  curl -X GET http://<host>:<port>/employees/all
  ```

### 2. Search Employee by ID

- **Endpoint:** `/{employeeId}`
- **Method:** GET
- **Description:** Searches for an employee by their ID.
- **Parameters:**
  - `employeeId` (path) - The ID of the employee to be searched.
- **Example Request:**

  ```
  curl -X GET http://<host>:<port>/employees/1
  ```

### 3. Filter Employees by Age

- **Endpoint:** `/greaterThanAge/{age}`
- **Method:** GET
- **Description:** Filters employees whose age is greater than the specified age.
- **Parameters:**
  - [`age`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A66%2C%22character%22%3A59%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (path) - The age to filter by.
- **Example Request:**

  ```
  curl -X GET http://<host>:<port>/employees/greaterThanAge/30
  ```

### 4. Filter Employees by City

- **Endpoint:** `/fromCity/{cityName}`
- **Method:** GET
- **Description:** Filters employees who are from the specified city.
- **Parameters:**
  - `cityName` (path) - The name of the city to filter by.
- **Example Request:**

  ```
  curl -X GET http://<host>:<port>/employees/fromCity/NewYork
  ```

### 5. Filter Employees by Age and City

- **Endpoint:** `/filterByAgeAndCity`
- **Method:** GET
- **Description:** Filters employees who are from a given city and whose age is greater than the specified age.
- **Parameters:**
  - `schoolId` (header) - School ID.
  - [`age`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A66%2C%22character%22%3A59%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (query) - Age of the employee.
  - [`cityName`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A81%2C%22character%22%3A33%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (query) - City of the employee.
- **Example Request:**

  ```
  curl -X GET http://<host>:<port>/employees/filterByAgeAndCity \
       -H "schoolId: S123" \
       -d "age=25&cityName=Chicago"
  ```

## Response Format

All endpoints return a JSON response with the relevant employee data or an error message in case of failure. Successful requests return a `200 OK` status code, while errors return a `500 Internal Server Error` with an error message.

### Success Response Example

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "age": 30,
    "city": "New York"
  }
]
```

### Error Response Example

```json
{
  "error": "An error occurred while processing your request."
}
```

Ensure to replace `<host>` and `<port>` with the actual values where your service is hosted.

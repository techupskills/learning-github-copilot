techupskills: @workspace what are the API's parameters and what does it return?

# Workspace:

Collecting workspace information

Here is a summary of the API endpoints, their parameters, and return types based on the implementation in `EmployeeController`:

### 1. **Get All Employees**
- **Endpoint**: `GET /employees/all`
- **Parameters**: None
- **Returns**: A list of all employees (`List<Employee>`).

---

### 2. **Search Employee by ID**
- **Endpoint**: `GET /employees/{employeeId}`
- **Parameters**:
  - `employeeId` (Path Variable, `Integer`): The ID of the employee to retrieve.
- **Returns**: An `Employee` object corresponding to the given `employeeId`.

---

### 3. **Filter Employees by Age**
- **Endpoint**: `GET /employees/greaterThanAge/{age}`
- **Parameters**:
  - `age` (Path Variable, `Integer`): The minimum age to filter employees.
- **Returns**: A list of employees (`List<Employee>`) whose age is greater than the specified value.

---

### 4. **Filter Employees by City**
- **Endpoint**: `GET /employees/fromCity/{cityName}`
- **Parameters**:
  - `cityName` (Path Variable, `String`): The name of the city to filter employees.
- **Returns**: A list of employees (`List<Employee>`) who are from the specified city.

---

### 5. **Filter Employees by Age and City**
- **Endpoint**: `GET /employees/filterByAgeAndCity`
- **Parameters**:
  - `schoolId` (Header, `String`): A required header parameter.
  - `age` (Query Parameter, `Integer`): The minimum age to filter employees.
  - `cityName` (Query Parameter, `String`): The name of the city to filter employees.
- **Returns**: A list of employees (`List<Employee>`) who are from the specified city and whose age is greater than the specified value.

---

### Notes:
- All endpoints return a `ResponseEntity` wrapping the result (e.g., `List<Employee>` or `Employee`) with appropriate HTTP status codes (`200 OK` for success, `500 Internal Server Error` for exceptions).

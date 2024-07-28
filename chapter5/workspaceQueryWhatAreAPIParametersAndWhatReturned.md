The [`EmployeeController`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A20%2C%22character%22%3A13%7D%5D "EmployeeController.java") class in [`EmployeeController.java`](EmployeeController.java) defines several API endpoints, each with its own set of parameters and return types. Here's a summary of the main endpoints, their parameters, and what they return:

1. **Get All Employees**
   - **Endpoint:** `/employees/all`
   - **Method:** GET
   - **Parameters:** None
   - **Returns:** A `ResponseEntity<Object>` containing a list of all [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") objects or an error message in case of failure.

2. **Search Employee by ID**
   - **Endpoint:** `/employees/{employeeId}`
   - **Method:** GET
   - **Parameters:**
     - [`employeeId`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A47%2C%22character%22%3A34%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (path variable, required): The ID of the employee to be viewed.
   - **Returns:** A `ResponseEntity<Object>` containing the [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") object with the specified ID or an error message in case of failure.

3. **Filter Employees by Age Greater Than Input**
   - **Endpoint:** `/employees/greaterThanAge/{age}`
   - **Method:** GET
   - **Parameters:**
     - [`age`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A65%2C%22character%22%3A59%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (path variable, required): The age to filter employees by.
   - **Returns:** A `ResponseEntity<Object>` containing a list of [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") objects whose age is greater than the input age or an error message in case of failure.

4. **Filter Employees by City**
   - **Endpoint:** `/employees/fromCity/{cityName}`
   - **Method:** GET
   - **Parameters:**
     - [`cityName`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A80%2C%22character%22%3A33%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (path variable, required): The city name to filter employees by.
   - **Returns:** A `ResponseEntity<Object>` containing a list of [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") objects from the specified city or an error message in case of failure.

5. **Filter Employees by Age and City**
   - **Endpoint:** `/employees/filterByAgeAndCity`
   - **Method:** GET
   - **Parameters:**
     - `schoolId` (header, required): School ID.
     - [`age`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A65%2C%22character%22%3A59%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (query parameter, required): Age of the employee.
     - [`cityName`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fcontroller%2FEmployeeController.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A80%2C%22character%22%3A33%7D%5D "src/main/java/com/hendisantika/springboot/swagger/controller/EmployeeController.java") (query parameter, required): City of the employee.
   - **Returns:** A `ResponseEntity<Object>` containing a list of [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") objects that match the age and city criteria or an error message in case of failure.

Each endpoint uses `ResponseEntity<Object>` to return either a list of [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") objects or a single [`Employee`](command:_github.copilot.openSymbolFromReferences?%5B%7B%22%24mid%22%3A1%2C%22path%22%3A%22%2Fworkspaces%2Fspring-boot-swagger%2Fsrc%2Fmain%2Fjava%2Fcom%2Fhendisantika%2Fspringboot%2Fswagger%2Fmodel%2FEmployee.java%22%2C%22scheme%22%3A%22file%22%7D%2C%7B%22line%22%3A16%2C%22character%22%3A13%7D%5D "src/main/java/com/hendisantika/springboot/swagger/model/Employee.java") object, depending on the endpoint, wrapped in an HTTP response with appropriate status codes. In case of an error, it returns an error message with an `HttpStatus.INTERNAL_SERVER_ERROR` status code.

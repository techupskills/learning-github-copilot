swagger: '2.0'
info:
  description: This API provides the capability to search Employee 
from an Employee Directory by different search parameters.
  version: 1.0.0
  title: Employee Directory API
  termsOfService: 'http://swagger.io/terms/'
  contact:
    email: support@example.com
  license:
    name: Apache 2.0
    url: 'http://www.apache.org/licenses/LICENSE-2.0.html'
host: 'localhost:8080'
basePath: '/employees'
tags:
  - name: employee
    description: Employee Management
schemes:
  - http
paths:
  /:
    get:
      tags:
        - employee
      summary: List all employees
      description: Returns a list of employees
      produces:
        - application/json
      responses:
        '200':
          description: successful operation
          schema:
            type: array
            items:
              $ref: '#/definitions/Employee'
  /{id}:
    get:
      tags:
        - employee
      summary: Find employee by ID
      description: Returns a single employee
      produces:
        - application/json
      parameters:
        - name: id
          in: path
          description: ID of employee to return
          required: true
          type: integer
          format: int64
      responses:
        '200':
          description: successful operation
          schema:
            $ref: '#/definitions/Employee'
        '400':
          description: Invalid ID supplied
        '404':
          description: Employee not found
definitions:
  Employee:
    type: object
    properties:
      id:
        type: integer
        format: int64
      name:
        type: string
      department:
        type: string
    required:
      - id
      - name

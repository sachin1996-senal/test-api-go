
## Overview
A simple Go HTTP API with one endpoint:


GET /hello-world?name=<name>

### Behavior
- If the first letter of `<name>` is **A–M** or **a–m**, respond with:
   json
  { "message": "Hello <name>" }

- If the first letter is **N–Z** or **n–z**, or if `<name>` is missing/empty, respond with:
    json
  { "error": "Invalid Input" }


## How to Run the Application
1. Make sure Go 1.18+ is installed (`go version`).
2. Run the server:
   powershell
   go run main.go

3. Test in a browser, Postman, or terminal:
 
   http://localhost:8080/hello-world?name=<your_name>

   Example:

   http://localhost:8080/hello-world?name=Alice


## How to Run the Tests
Run unit tests:
powershell
go test -v


## Assumptions
- Only English letters (A–Z, a–z) are checked.  
- Case-insensitive comparison (e.g., `alice` and `Alice` both valid).  
- Empty or whitespace names are invalid.  
- Invalid inputs return HTTP **400** with `{ "error": "Invalid Input" }`.


Example dynamic responses:
| Request | Response | Status |
|----------|-----------|--------|
| `/hello-world?name=Alice` | `{ "message": "Hello Alice" }` | 200 |
| `/hello-world?name=Zane`  | `{ "error": "Invalid Input" }` | 400 |
| `/hello-world`            | `{ "error": "Invalid Input" }` | 400 |

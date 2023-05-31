# gotaskmanager
Task manager built with Go

### Features
1. User register and login
2. User can add tasks
3. User can write notes associated with each tasks.
4. Authentication and Authorization using JWT

### Endpoints
| URL | HTTP Verb | functionality |
| -------- | --------- | ------ |
| '/user/login' | POST | user login with email and password |
| '/user/register' | POST | user register |
| '/task/{id}' | GET | user register |
| '/task' | POST | user create tasks |
| '/task/{id}' | PUT | user update tasks |
| '/task{id}' | DELETE | user delete tasks |
| '/tasks' | GET | fetch all tasks |


### Tools
1. Go
2. MongoDB



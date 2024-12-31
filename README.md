# todoBackendGokit
Backend implementation of Todo app using Golang, Go-kit framework and Postgresql 

Deployed at https://todobackendgokit.onrender.com

Example- 
GET - 
https://todobackendgokit.onrender.com/tasks/get/id
Response - {"id":"e24c1fde-98f4-4858-bba8-5f51fed62823","title":"Learn Go-kit","status":"pending"}

GET ALL - 
https://todobackendgokit.onrender.com/tasks/getall
Response - [{...},{...}]

CREATE - 
https://todobackendgokit.onrender.com/tasks/create
Body - { "title": "Learn Go-kit", "status": "pending" }

UPDATE - 
https://todobackendgokit.onrender.com/tasks/update
Body - {"id":"e24c1fde-98f4-4858-bba8-5f51fed62823","title":"leran golang","status":"completed"}

DELETE - 
https://todobackendgokit.onrender.com/tasks/delete/id

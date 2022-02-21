# todo-api
Api for todos.

 - Uses PostgreSQL as database 
 - Caches with Redis
 
 ### Running  🚀
 
 In the main directory

    docker-compose up

### Endpoints📍

Postman collection are in the main directory

**GET /api/todos/:guid**

Gets todo

**UPDATE /api/todos/:guid** 

Marks todo as done

**DELETE /api/todos/:guid**

Deletes todo

**POST /api/todos**

Creates todo	
	
    title: 		 string
    description: string

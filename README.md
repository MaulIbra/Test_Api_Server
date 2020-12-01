# Documentation

![db_user_new](https://user-images.githubusercontent.com/46882131/100691378-2950e380-33bb-11eb-8df6-6bb209bd5eed.PNG)

**Setup Configuration**
- Open config.env and change all environment variable

**Step Running**
1. Import .sql to your local machine
2. Running main.go with command go run main.go

**API ROUTES**
- **Authentication** (to get token)
1. /account/login [POST]
2. /account/register [POST]
- **User**
1. /user [POST]
2. /user/{page}/{limit} [GET] 
3. /user/{id} [GET]
4. /user/{id} [PUT]
5. /user/{id} [DELETE]

**Documentation Postman**
- https://documenter.getpostman.com/view/11868527/TVmJiK8X

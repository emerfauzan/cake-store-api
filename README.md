## Intro
Welcome to Cake Store API. This API allows you to manage cake data. You can add, edit, show and delete your own cake!

##Features
This API contains:
1. Login API (POST /auth/login)
    This API is used to login and get token. The token will be used to the next API.
2. Get All Cakes (GET /cakes)
    This API is used to show all active cakes data.
3. Get Cake by Id (GET /cakes/{id})
    This API is used to show a cake data.
4. Create Cake (POST /cakes)
    This API is used to add a cake data.
5. Update Cake (PUT /cakes/{id})
    This API is used to update a cake data by id.
6. Delete Cake (DELETE /cakes/{id})
    This API is used to delete a cake data by id.

The API documentation is provided in folder documentation

## How to Run Project
1. After clone this project, create your cake database
2. Copy .env.sample to .env
3. Change the database configuration in .env file
4. RUN "make migrate-up"
5. RUN "make seed"
6. RUN "make run"


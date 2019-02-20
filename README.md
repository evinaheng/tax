# Bills Tax Lot App

## How To Develop

## Cautions
- `vendor` folder included for now, should be included on gitignore
- `api-bills` executeable script app also included for now, should be included on gitignore
- included for now in case you cannot dep update regarding the library
- for full link git in `https://github.com/evinaheng/tax`
 
### Onboarding
- Install Golang `1.11.2` & Golang Dep `0.5`
- This implementation is inspire from clean architechture concept for Golang Base Project
- Base Project for rest api is consist 4 level such as `delivery`, `usecase`, `controller` and `repository` (internal/repository)
- `repository` will be contain all kind functional that related to internal library such as httpClient, DB call, cache etc
- `controller` will be contain all kind functional that contain multiple repository and serve it as one function, also will contain business logic (internal/controller)
- `usecase` will be works as one big functional (goals) on one of endpoint service (internal/usecase)
- `delivery` will be initialize all needed packages and serve it on the service routes (api/tax/)
- `entity` will be contain all kind of struct object that using on the project <- mostly representative from db structure

### Compiling
- For UNIX, build using `make build`

### Running
- Run docker, refer to [this document](../docker/README.md)


### For Running the Application
- Copy `files/etc/tax/development` to `development` to main folder
- use dep to get all related library `dep ensure -v`
- Type `make build` to make sure all being compile
- Type `./api-tax` to start the application


### ALL Endpoint
- `http://localhost:9006/v1/tax/getBills` to get all list bills 
- `localhost:9006/v1/tax/createBills?name=BigBurger&price=1040&type_code=1` example to create new bills record
- please access local file on the frontend.html to get sample view from getBills

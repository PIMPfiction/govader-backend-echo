# Govader Backend with Echo


Govader-Backend is based on GoVader Package[https://github.com/jonreiter/govader](https://github.com/jonreiter/govader)

Main Package Repository is Govader-Backend(https://github.com/PIMPfiction/govader_backend)

## Usage:

### Clone this Repo and run
```sh
git clone https://github.com/pimfiction/govader-backend-sample
cd govader-backend-sample
go run src/main.go
```

### Deploy on Heroku
Download heroku cli from official stie (https://devcenter.heroku.com/articles/heroku-cli)
```sh
git clone https://github.com/pimfiction/govader-backend-sample
cd govader-backend-sample
heroku login # will open browser for you to login heroku
heroku create 
git push heroku main
heroku open # will open the browser to your app's URL directly
```


### Sample Get Request

#### GET: http://localhost:8080?text=I%20am%20looking%20good

### Sample Post Request 

#### POST: http://localhost:8080/
### RequestBody: ```{"text": "I am looking good"}```
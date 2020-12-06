# charlie-parker 🎷🚗
*Making parking smoother than the jazz king, [Yardbird](https://en.wikipedia.org/wiki/Charlie_Parker), himself!* This app gave me an opportunity to flex and sharpen my Go/REST api/AWS Dynamo DB/Docker skills!

## Project Structure
```
charlie-parker
 ├── docker-compose.local.yml
 ├── Dockerfile
 ├── go.mod
 ├── go.sum
 ├── main.go -- blank main file to allow for testing
 ├── README.md
 ├── tools.sh -- command line tools for running the app
 ├── cmd
 |    ├── seeder
 |    |    └── main.go -- app entry for seeding local dynamo
 |    └── server 
 |         └── main.go -- app entry for running the api server
 ├── internal
 |    ├── config
 |    |    └── config.go -- init() for app-wide configuration
 |    ├── helpers
 |    |    ├── rates.go         -- helper funcs for routes in \routes\rates.go
 |    |    ├── routemetrics.go  -- helper funcs for route metrics and routes in \routes\routemetrics.go
 |    |    ├── util_test.go     -- tests for util.go
 |    |    ├── util.go          -- general helper functions for data manipulation
 |    |    ├── validate_test.go -- tests for validate.go
 |    |    └── validate.go      -- validation functions for route inputs
 |    ├── routes
 |    |    ├── rates.go        -- rate-related route handlers
 |    |    └── routemetrics.go -- metrics-related route handlers
 |    ├── seeder
 |    |    ├── seed_data.go -- defines a list of CreateRateInput used to seed
 |    |    └── seeder.go    -- exports Run() that runs the seeder
 |    └── server
 |         └── server.go    -- exports Start() that starts the server
 ├── pkg \ types
 |    ├── rates.go        -- defines the rate struct and input/output types to rate-related routes
 |    └── routemetrics.go -- defines the route metrics struct and input/output types to metrics-related routes
 |    └── utiltypes.go    -- defines the BaseOutput type that contains Ok and Error fields
 ├── utils
 |    └── utilroutes.go -- HeartbeatRoute() to check app alive-ness
 └── vendor -- vendored dependencies     
```

## Functionality
This app allows for the storage and retrieval of [parking rates](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/pkg/types/rates.go#L4) that have a comma separated list of days for which they cover, a time span in the format "HHMM-HHMM", a time zone, and a price. Rates must not define a time range that is already covered by a rate. For example, a rate that covers 9am-2pm on Fridays in the timezone America/Chicago may not be created if a rate that covers 12pm-1pm on Fridays in the timezone America/Chicago already exists.

A set of [start and end times](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/pkg/types/rates.go#L45) strings may be sent to the app's server, if there is a rate that covers that time range, its price will be returned to the client. [Start and end must](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/helpers/validate.go#L162):

  1. parse in [ISO-8601](https://en.wikipedia.org/wiki/ISO_8601) format
  2. be in the same year.
  3. be on the same day. 
  4. not be the same time. 
  5. be one earlier hour and one later hour.
  6. be in the same timezone.

In order for a set of start and end times [to match to an existing rate](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/helpers/util.go#L160), a rate must exist:

  1. that covers the day on which start and end fall.
  2. that matches the start and end's timezone.
  3. for which start is greater than or equal to its start time
  4. for which start is less than its end time
  5. for which end is greater than its start time
  6. for which end is less than or equal to its end time
  
Beyond this business functionality, average response time, API endpoint hits, number of successful exchanges, and number of failed exchanges are metrics measured for each route defined for the server.

## Building Docker Containers
Currently, this app is only set up for a local environment. You will need Docker and Go installed in order to run. (You can either follow the instructions below or alternatively run the file [tools.sh](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/tools.sh#L5) with the argument `local` to run a script to build and test this app entirely.)

In your command prompt navigate to where you have cloned this repo:

> Run Dynamo on port 8000: `docker run -d -p 8000:8000 -v ~/:/var/lib/dynamodb -it --rm --name cp-dynamo instructure/dynamo-local-admin`

> Build the seeder: `docker build -t charlie-parker-seeder:latest --build-arg app=seeder .`

> Build the server: `docker build -t charlie-parker-server:latest --build-arg app=server .`

> Run the seeder: `docker run -e AWS_ACCESS_KEY_ID=key -e AWS_SECRET_ACCESS_KEY=secret -e AWS_REGION=us-east-1 --name cp-seeder charlie-parker-seeder:latest`

> Run the server: `docker run -e AWS_ACCESS_KEY_ID=key -e AWS_SECRET_ACCESS_KEY=secret -e AWS_REGION=us-east-1 -p 8554:8554 --name cp-server charlie-parker-server:latest`

_(The environment variables related to AWS are needed to [connect to the local dynamo tables](https://hub.docker.com/r/instructure/dynamo-local-admin))_ 

## Business Logic Testing
Tests are defined for two files: internal\helpers\\[utils.go](https://github.com/noahwill/charlie-parker/blob/master/internal/helpers/util.go) and internal\helpers\\[validate.go](https://github.com/noahwill/charlie-parker/blob/master/internal/helpers/validate.go) in [utils_test.go](https://github.com/noahwill/charlie-parker/blob/master/internal/helpers/util_test.go) and [validate_test.go](https://github.com/noahwill/charlie-parker/blob/master/internal/helpers/validate_test.go) respectively. These two files contain most of the business logic and do not need a DB connection nor an HTTP request to test. These are the tests run [when Docker is building](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/Dockerfile#L6) the app; they may also be run individually.

## Route Testing
In order to test the routes defined for the server, build the app and use the following curl commands.

### GET all Rates
[This](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/server/server.go#L31) [route](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/routes/rates.go#L17) gets all rates that are in the rates table. If you did not run the seeder or have deleted all of the rates, none will be returned.

> Mac/Linux/Windows: `curl -X GET http://localhost:8554/api/v1/rates`

### POST to create a rate
[This](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/server/server.go#L32) [route](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/routes/rates.go#L40) creates a rate based on the following required input:
  - `Days` any substring of `"sun,mon,tues,wed,thurs,fri,sat"`
  - `Times` a string range of 24-hour time hours and minutes in the format of `"HHMM-HHMM"`
  - `TZ` a string timezone (i.e. `"America/Chicago"`)
  - `Price` an integer (represents number of cents charged per hour)

> Mac/Linux: `curl -X POST -H "Content-Type: application/json" -d '{"Days": "fri", "Times": "1600-1800", "TZ": "America/Chicago", "Price": 1800}' http://localhost:8554/api/v1/rates/create`

> Windows: `curl -X POST -H "Content-Type: application/json" -d "{\"Days\": \"fri\", \"Times\": \"1600-1800\", \"TZ\": \"America/Chicago\", \"Price\": 1800}" http://localhost:8554/api/v1/rates/create`

### POST to overwrite all the routes
[This](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/server/server.go#L33) [route](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/routes/rates.go#L71) is a useful route for batch creating a large set of new rates. It overwrites all existing rates in the DB. This is obviously not a useful route if this were a real world app where we'd probably want to keep old ratese around, but for now, since this is all local and the containers will be torn down anyway, this is a useful route incase the user would like to test creating a whole bunch of different rates. The required input is a list of inputs of the same fields used in the create rate route:

> Mac/Linux: `curl -X POST -H "Content-Type: application/json" -d '{"Rates": [{"Days": "fri", "Times": "1600-1800", "TZ": "America/Chicago", "Price": 1800}, {"Days": "fri", "Times": "0900-1200", "TZ": "America/Chicago", "Price": 500}]}' http://localhost:8554/api/v1/rates/update/all`

> Windows: `curl -X POST -H "Content-Type: application/json" -d "{\"Rates\": [{\"Days\": \"fri\", \"Times\": \"1600-1800\", \"TZ\": \"America/Chicago\", \"Price\": 1800}, {\"Days\": \"fri\", \"Times\": \"0900-1200\", \"TZ\": \"America/Chicago\", \"Price\": 500}]}" http://localhost:8554/api/v1/rates/update/all`

### POST to get the price for a timespan
[This](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/server/server.go#L35) [route](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/routes/rates.go#L102) tries to find a rate based on the following required input:
  - `Start` a string in the format `"2017-01-06T17:00:00-06:00"`
  - `End` a string in the format `"2017-01-06T018:00:00-06:00"`

_(These will return a price only if you've used the create or overwrite examples above)_
> Mac/Linux: `curl -X POST -H "Content-Type: application/json" -d '{"Start": "2017-01-06T17:00:00-06:00", "End": "2017-01-06T18:00:00-06:00"}' http://localhost:8554/api/v1/park`

> Windows: `curl -X POST -H "Content-Type: application/json" -d "{\"Start\": \"2017-01-06T17:00:00-06:00\", \"End\": \"2017-01-06T18:00:00-06:00\"}" http://localhost:8554/api/v1/park`

### GET route metrics
[This](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/server/server.go#L39) [route](https://github.com/noahwill/charlie-parker/blob/cd87ad3e2221173035476941f95c314046cb8cdd/internal/routes/routemetrics.go#L17) gets the route metrics for all of the routes defined by this app. _If hit right after a fresh build of the app, this route will return no metrics. Hit a few other routes, or this one a couple more times, then call this one to see the metrics come in!_

> Mac/Linux/Windows: `curl -X GET http://localhost:8554/api/health/routes`

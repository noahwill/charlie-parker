# charlie-parker 🎷🚗
*Making parking smoother than the jazz king, [Yardbird](https://en.wikipedia.org/wiki/Charlie_Parker), himself!*

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


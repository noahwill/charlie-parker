# charlie-parker 🎷🚗
*Making parking smoother than the jazz king, [Yardbird](https://en.wikipedia.org/wiki/Charlie_Parker), himself!*

## Project Structure
```
go.mod
go.sum
Dockerfile
├── cmd
|    ├── seeder
|    |    └── main.go -- app entry for seeding local dynamo
|    └── server 
|         └── main.go -- app entry for running the api server
├── internal
|    ├── config
|    |    └── config.go -- init() for app-wide configuration
|    ├── helpers
|    |    ├── rates.go    -- helper funcs for routes in \routes\rates.go
|    |    ├── util.go     -- general helper functions for data manipulation
|    |    └── validate.go -- validation functions for route inputs
|    ├── routes
|    |    └── rates.go -- rate-related route handlers
|    ├── seeder
|    |    ├── seed_data.go -- defines a list of CreateRateInput used to seed
|    |    └── seeder.go    -- exports Run() that runs the seeder
|    └── server
|         └── server.go    -- exports Start() that starts the server
├── pkg \ types
|    ├── rates.go     -- defines the rate struct and input/output types to rate-related routes
|    └── utiltypes.go -- defines the BaseOutput type that contains Ok and Error fields
├── utils
|    └── utilroutes.go -- HeartbeatRoute() to check app alive-ness
└── vendor -- vendored dependencies     
```


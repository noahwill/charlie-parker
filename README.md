# charlie-parker 🎷🚗
*Making parking smoother than the jazz king, [Yardbird](https://en.wikipedia.org/wiki/Charlie_Parker), himself!*

## Project Structure
```
├── cmd
|    ├── seeder
|    |    └── main.go -- app entry for seeding local dynamo
|    └── server 
|         └── main.go -- app entry for running the api server
├── internal
|    ├── seeder
|    |    ├── seed_data.go
|    |    └── seeder.go    -- exports Run() that runs the seeder
|    └── seeder
|         └── seeder.go    -- exports Start() that starts the server
└── pkg \ types
     └── rates.go -- defines the rate struct
```


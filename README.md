# Go CRUD 

A backend application to perform Create and Read operations for tools/resources for frontend/fullstack developers and UX/UI designers. 

## Note
View the older repository using this [Link](https://github.com/DanielekPark/gosql)

## Languages used
- Golang/Go
- SQL

### Database platform
[PlanetScale](https://planetscale.com/) MySQL database

### Go packages

[Fiber](https://docs.gofiber.io/v1.x/) Installing Fiber, a Express inspired web framework written in Go
```sh
go get -u github.com/gofiber/fiber
```

[MySQL Driver](https://github.com/go-sql-driver/mysql?tab=readme-ov-file#installation) Enables the use of SQL using Go
```sh
go get -u github.com/go-sql-driver/mysql
```

[Air](https://github.com/cosmtrek/air) For running live reloads
```sh
go install github.com/cosmtrek/air@latest
```

[Dotenv](https://github.com/joho/godotenv) To keep environment variables safe
```sh
go get github.com/joho/godotenv
```

### Command line

[Link](https://stackoverflow.com/questions/43983718/how-can-i-globally-set-the-path-environment-variable-in-vs-code) to the directions to make changes to VS Code
```sh
go env GOPATH
```
```sh
export PATH="$HOME/go/bin:$PATH"
```


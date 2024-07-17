# Domain Driven Design with Go and PostgreSQL

![Go](https://img.shields.io/badge/Go-1.16-blue) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13-blue) ![DDD](https://img.shields.io/badge/DDD-Hexagonal%20Architecture-red)

## Introduction

Welcome to my project where I am currently learning and implementing Domain Driven Design (DDD), also known as Hexagonal Architecture, using Go and PostgreSQL. This repository serves as a playground for experimenting with DDD concepts and building a robust application architecture.

## Features

- **Domain Driven Design:** Emphasis on the core domain and domain logic.
- **Hexagonal Architecture:** Decoupled application components making them easier to test and maintain.
- **Go Language:** Using Go for its simplicity and performance.
- **PostgreSQL:** Robust and reliable relational database.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [PostgreSQL](https://www.postgresql.org/download/) (version 13 or later)

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/Nezent/GoChi.git
    cd GoChi
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

<!-- 3. Set up the database:
    ```sh
    psql -U postgres -c "CREATE DATABASE ddd_hexagonal;"
    ``` -->

3. Run the application:
    ```sh
    make run
    ```

## ToDo

- [x] Set up initial project structure
- [x] Implement domain models
- [x] Create repository interfaces
- [x] Set up PostgreSQL database connection
- [x] Implement application services
- [x] Create REST API endpoints using GoChi
- [ ] Write unit and integration tests
- [ ] Add documentation for each module
- [ ] Deploy the application

## Learning Resources

- [Domain-Driven Design Reference](https://www.domainlanguage.com/ddd/reference/)
- [Hexagonal Architecture Overview](https://alistair.cockburn.us/hexagonal-architecture/)
- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

<!-- ## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

--- -->

Happy coding! 🚀

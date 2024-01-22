# HotelManagementSystem

## Introduction
HotelManagementSystem is a microservices-based hotel management system designed for efficient hotel operations. This system enables functionalities such as hotel creation, removal, management of contact details, listing officials, and generating detailed reports.

## Features
- Hotel and contact info all CRUD operations.
- Generation of detailed reports via Message Queues.

## Technologies Used
- Go
- C#
- RabbitMQ
- PostgreSQL
- Docker
- Entity Framework Core

## Installation and Setup
To set up the HotelManagementSystem, follow these steps:
1. Clone the repository: `git clone www.github.com/AlpTalhaYazar/HotelManagementSystem.git`.
2. Navigate to the project directory.
3. Navigate to the HotelManagementService directory.
4. Run the following command: `go run src/main.go`.
5. Before running dotnet services, first run the following command: `docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management` to run RabbitMQ.
6. Now you can run the dotnet services from your IDE or from the command line using `dotnet run` command.

## Usage
Here's how to use the system:
1. After running the system, you can access the API documentation from repository root directory by opening the Postman collection file.
2. There you will see rest endpoints for the hotels, contact infos and reports.

## Contributing
See the GitHub documentation on [creating a pull request](https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request).

## License
[Specify the license or 'This project is licensed under the [License Name] - see the LICENSE.md file for details'].

## Contact
For further queries or suggestions, please contact Alp Talha Yazar at my mail address [Send Mail](mailto:AlpTalhaYazar@gmail.com) or you can contact me from [LinkedIn](https://www.linkedin.com/in/alptalhayazar/).
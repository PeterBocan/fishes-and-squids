# Ocean Orchestra Full Stack Interview Challenge

Welcome to the Ocean Orchestra Full Stack Interview Challenge! This repository contains a simple full-stack application designed to test your skills in both frontend (React/ TypeScript) and backend (Golang) development.

We do not want you to spend more than **2 hours** on this challenge. Focus on the core goals first and if you have time, feel free to explore any optional goals.

## Goals

## Core Goals

Our core goals are to see how you go about building a simple full stack application that pulls data from a database and displays it to a user.

1. **Seed the database** - Seed a PostgreSQL database with sample data from `./db/products.json`
2. **GET /products** - Implement a RESTful API in Go that serves the list of data from the database `./api`
3. **List data on the frontend** - Create a React frontend that calls the Go API and displays the data in a list `./web` using some basic styling

### Optional Goals

If you have time, feel free to explore any of the following optional goals. These are not in a priority order, we're interested in seeing what you choose to prioritise and how you go about implementing it. Think about what you'd like to show us about your skills and approach. In our follow up interview, we will ask you about your choices and thought process and how you might go about any of the optional goals you didn't get to.

- **Create a new product** - Let the user create a new product via a form in the React frontend that sends a POST request to the Go API to add the product to the database. (Let the user set the name field and then provide dummy data for the other fields).
- **Delete a product** - Let the user delete a product from the list in the React frontend that sends a DELETE request to the Go API to remove the product from the database.
- **Update a product** - Let the user update a product's name in the list in the React frontend that sends a PUT/PATCH request to the Go API to update the product in the database.
- **Frontend sorting** - Let the user sort the list of products by name or price in ascending or descending order.
- **Frontend testing** - Add some tests to the React frontend. It doesn't need to be exhaustive, just enough to show you understand the concepts and best practices.
- **Backend testing** - Add some tests to the Go backend. It doesn't need to be exhaustive, just enough to show you understand the concepts and best practices.
- **Loading/ Error Handling** - Add loading states and error handling to the React frontend when calling the Go API.
- **Backend filtering** - Let the user filter the list of products by a search term in the React frontend that sends a query parameter to the Go API to filter the products in the database.
- **Improved styling** - Improve the styling of the React frontend to make it look more polished and user-friendly.

## Getting Started

### Prerequisites

Ensure you have the following installed on your machine:

- Go
- Node.js
- Docker and Docker Compose

### Project Installation

1. `make install` / `cd ./web && yarn install` - Install dependencies for the React frontend.
2. `make up` / `docker-compose up -d` - Builds and starts the PostgreSQL database using Docker Compose.

### Running Everything

1. `make api` / `cd api && go run .` - Start the Go backend server.
2. `make web` / `cd web && yarn start` - Start the React frontend (see logs for the port it's running on).
3. `make up` / `docker-compose up -d` - Builds and starts the PostgreSQL database using Docker Compose.

You can optionally run `make db` / `docker compose logs -f db` to see the database logs.

You can run `make down` / `docker compose down -v` to stop and remove the database container.

### Project Structure

- `./db`: Contains database initialization scripts and sample data. `./db/init.sql` gets run on database startup, and `./db/products.json` is sample data.
- `./api`: Contains the Go backend code.
- `./web`: Contains the React frontend code.
- `docker-compose.yml`: Docker Compose configuration file to set up the PostgreSQL database.
- `Makefile`: Contains commands to manage the project.

## Submission

1. Please zip your project folder and email it back to us or share a link to a public Git repository.
2. Please let us know how long you spent on the challenge and any areas you would have liked to improve or add if you had more time.
3. Feel free to use AI tools as you normally would, and give us an overview of how you used them. We expect you to understand and be able to justify every line of code without AI assistance. So only include code you understand.

## Tips & What We're Looking For

- You could seed data via `./db/init.sql` or programmatically via Go. Choose whatever you prefer.
- Use any libraries or frameworks you like and feel free to change the config/ project structure as needed.
- We will be running `make up`, `make api` and `make web` to test your solution. So ensure these commands work as expected.
- Add comments everywhere you think it's necessary to explain your thought process. There's more work than time on this project, we don't want you to complete everything, but we want to understand how you think, work and decide what to prioritise. Also adding comments for best practices that you're not implementing because of time such as error handling, validation, security, testing etc is a great way to show you understand these areas even if you don't have time to implement them.
- You can use any styling framework you like, we've installed tailwind for you already, which we use. But anything is fine. This is not a design challenge but we do want to see how you approach styling and layout.
- Feel free to initialise the project as a git repo and make commits as you go. This helps us see how you work and think.

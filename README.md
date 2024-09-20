# PointaFam

## Introduction

**PointaFam** is a revolutionary platform designed to connect local farmers with urban retailers, streamlining the distribution of fresh, locally-sourced produce. By reducing the friction between supply and demand, PointaFam helps retailers discover local farms and source high-quality products at fair prices. This platform is built with a user-centric approach, focusing on ease of use and reliable functionality.

- **Backend**: Golang, GORM, PostgreSQL
- **Frontend**: HTML, CSS, JavaScript, Bootstrap

### Deployed Site
Visit the live version of PointaFam here: [PointaFam Deployed Site](https:pointafam.onrender.com)

### Final Project Blog Article
Read more about the journey and development process behind PointaFam in the final blog post: [Final Blog Article](#)

### Author's LinkedIn
Connect with me on LinkedIn: [LinkedIn](https://www.linkedin.com/in/ferdynand-odhiambo)

---

## Installation

To run this project locally, follow the steps below:

### Prerequisites

Ensure you have the following installed on your machine:
- [Go](https://golang.org/doc/install) (version 1.19+)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Node.js](https://nodejs.org/en/download/) (for frontend build)
- Git

### Clone the Repository

```bash
git clone https://github.com/MeFerdi/pointafam_capstone.git
cd PointaFam
```
### Backend Setup

#### Install Go dependencies:

```bash

go mod download
```
#### Set up PostgreSQL:

- Create a new PostgreSQL database for PointaFam.
- Configure the .env file with your database credentials. Example:

```bash
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=pointafam_db
```
- Run database migrations:

```bash

go run migrations/migrate.go
```
Start the Go server:

```bash

go run main.go
```
The backend server will start on http://localhost:8080.

#### Frontend Setup

- Navigate to the frontend directory:

```bash

cd frontend
```
- Install frontend dependencies:

```bash

npm install
```
#### Build the frontend:

```bash

npm run build
```
Serve the frontend: You can use any static server or integrate it with the backend.

#### Usage
##### For Retailers:

- Browse farms: Retailers can browse local farms and discover produce based on their location and preferences.
- Filter produce: Search for specific products and filter them by categories like organic, price, and availability.
- Add to cart: Once you find the desired products, you can add them to your cart and proceed to checkout.

##### For Farmers:

- Register your farm: Farmers can create an account and register their farms, adding details like location, product offerings, and availability.
- Update offerings: Keep your product list updated with available produce, pricing, and seasonal offerings.
- Manage orders: Handle incoming orders from retailers and ensure timely delivery.

#### Developer Use

##### To test the project locally:

- Ensure both backend and frontend servers are running.
- Access the app by visiting http://localhost:8080 for backend and http://localhost:3000 for frontend.

#### Contributing

- Contributions are welcome! Follow these steps to contribute:

##### Fork the repository:
- Click the "Fork" button at the top of this repository.
- This will create a copy of this repository under your GitHub account.

- Clone the forked repository:

```bash

git clone https://github.com/yourusername/pointafam_capstone.git
cd PointaFam
```

- Create a new branch:

```bash

git checkout -b feature/your-feature-name
```
- Make changes and commit:

- Make your code changes.
- Add tests where applicable.
- Run the tests to ensure everything works as expected.
- Commit your changes:

```bash
git add .
git commit -m "Add feature: your feature description"
```
Push your changes:

```bash

git push origin feature/your-feature-name
```
- Open a pull request:

- Go to the original repository on GitHub.
- Click the "New Pull Request" button.
- Select the branch you just pushed to and submit the pull request.

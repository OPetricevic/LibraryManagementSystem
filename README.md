# Library Management System

Welcome to the Library Management System! This system allows you to manage your library's collection of books and patrons.

## Getting Started

Before you can use the system, you'll need to set up your own database and configure the application to use it. Follow these steps:

### Prerequisites

- Go: Make sure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

- MySQL: You need a MySQL database server. You can download and install it from [here](https://dev.mysql.com/downloads/installer/).

### Configuration

1. Clone the repository:

git clone https://github.com/YourUsername/LibraryManagementSystem.git
cd LibraryManagementSystem


2. Create a `.env` file in the root of the project and add your MySQL database credentials:

- DB_PASSWORD=your_password_here

3. Replace `your_password_here` with your MySQL database password.

4. Install project dependencies:

- go mod tidy


5. Build the application:

- go build -o LibraryManagementSystem

6. Run the application:

- ./LibraryManagementSystem


7. If the application successfully connects to your MySQL database, you'll see "Database is reachable!" in the console.

### Usage

You can now use the Library Management System to manage your library's collection. The system allows you to add books, patrons, and perform various library management tasks.

### Troubleshooting

If you encounter any issues during setup or usage, please check the following:

- Ensure your MySQL server is running.
- Make sure you have Go and project dependencies installed.






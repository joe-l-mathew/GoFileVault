# GoFileVault

GoFileVault is a Golang project that transforms your old laptop into a self-hosted cloud server. Users can create accounts using their email and password, securely upload files, and manage their data. The application utilizes Gorilla Mux as the router and Gorm as the ORM, powered by SQLite for persistent storage.

**Key Features:**

- **User Authentication:** Create accounts with email and password for secure access, using JWT for authentication
- **File Upload and Storage:** Upload and securely store files on your personal server.
- **SQLite Database:** User upload details are persistently stored for efficient data management.
- **Query and Download:** Easily query and download your uploaded files.

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/GoFileVault.git
   cd GoFileVault
   ```

 ## Configure the Application:

 - Update the database configurations in config.go.
 - Customize other settings in the configuration files as needed.


## Build and Run:
```bash
go build
./GoFileVault
```

## Access the Application:

 - Open your web browser and go to http://localhost:8080 to start using GoFileVault.

## Contributing
 - Contributions are welcome! Fork the repository, create a new branch, and submit a pull request with your enhancements.

## License
 - This project is licensed under the MIT License.
- Feel free to modify or expand upon it according to your specific project details.

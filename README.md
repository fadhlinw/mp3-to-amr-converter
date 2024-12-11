# MP3 to AMR ConvertApp with FFmpeg Integration

This project integrates FFmpeg with a Go application to process audio files.

## Prerequisites

1. Install Docker on your OS or server. Follow the official Docker installation guide: [Docker Documentation](https://docs.docker.com/get-docker/).

## Steps to Run

1. **Build the Docker Image**  
   Run the following command in the project directory:
   ```bash
   docker build -t convertapp .
   ```

2. **Prepare Input Files**  
   Ensure the input files are present in the same repository as the application.

3. **Run the Application**  
   Use the following command to run the Go application:
   ```bash
   go run main.go
   ```

## Notes

- The application expects audio files for processing to be available in the project directory.
- Docker is used for building the environment with FFmpeg integration.

## License

MIT


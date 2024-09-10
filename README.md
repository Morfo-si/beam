# BEAM: Bot Engine for Application Messaging

**BEAM** (Bot Engine for Application Messaging) is a prototype RESTful API designed to facilitate communication between applications and Large Language Models (LLMs). It allows users to send custom prompts and questions to various LLMs and retrieve their responses for display in a web interface application.

## Getting Started

### Prerequisites

To run BEAM, you'll need:

- [Go](https://golang.org/doc/install) installed on your machine.
- [Git](https://git-scm.com/) to clone the repository.
- [Ollama](https://ollama.com) OR [OpenShiftLightSpeed](https://github.com/openshift/lightspeed-service) running locally.
- A `.env` file for configuration.

### Setup Instructions

1. **Clone the repository**:

   ```bash
   git clone https://github.com/your-username/beam.git
   cd beam
   ```

2. **Copy the `.env.example` file**:
   The project includes an example environment file, `.env.example`. To configure your environment, copy this file and rename it to `.env`:

   ```bash
   cp .env.example .env
   ```

3. **Configure the `.env` file**:
   Open the `.env` file in a text editor and update the necessary environment variables (such as `HOST`, `PORT`, `LLM_NAME`, `LLM_API`, and `SSLMODE`) to match your local or production environment setup.

   Example `.env` file:

   ```bash
   HOST=localhost
   PORT=8081
   ENGINE=beam
   LLM_NAME=llama3:latest
   LLM_API=http://localhost:11434/api/generate
   ```

### Using the Makefile

This project includes a `Makefile` to simplify common tasks.

- **Run the application**:
  
  Use the `run` rule to start the application. This will load the environment variables from the `.env` file and execute the Go application.
  
  ```bash
  make run
  ```

- **Build the application**:
  
  To compile the project and create an executable binary:
  
  ```bash
  make build
  ```

- **Test the application**:
  
  Run tests for the project:
  
  ```bash
  make test
  ```

- **Clean the project**:
  
  Use the `clean` rule to remove compiled binaries and cached files:
  
  ```bash
  make clean
  ```

- **Install dependencies**:
  
  Install Go dependencies specified in the project:
  
  ```bash
  make deps
  ```

### Disclaimer

**BEAM** is a prototype created for experimentation and testing purposes only. It is not intended for production use and is not a commercial solution. The project is not backed or supported by my employer, and it comes with no warranties or guarantees. Use at your own risk.

---

If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request.

### Key Sections

- **Getting Started**: Guides users through cloning the repository and setting up the `.env` file.
- **Using the Makefile**: Provides commands to run, build, test, and clean the project using the included `Makefile`.
- **Disclaimer**: Clarifies that this project is a prototype and not a commercial solution supported by your employer.

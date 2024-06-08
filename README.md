# MailFort
MailFort is an email service project that allows you to send confirmation emails and password reset emails. It provides flexibility in configuration and supports both gRPC and REST API for interaction.

## Getting Started
To run the MailFort service, follow these steps:

1. **Fill out the configuration file:** Create a configuration file (e.g., `local.yaml`) and provide the necessary configurations for the MailFort service. You can find an example configuration file in the `config` directory.

2. **Export the configuration path:** Set the environment variable `CONFIG_PATH` to point to your configuration file.
```bash
export CONFIG_PATH=./config/local.yaml
```

3. **Add credential for mail client:** Create file `.env` and write parameters.

4. **Run in development mode:** Start the MailFort service in development mode using the task command:
```bash
task dev
```

## Interaction
MailFort supports interaction through both gRPC and REST API.

### gRPC Interaction
To interact with MailFort using gRPC, you can use the generated gRPC client. Below is an example in Go:

```go
// Example gRPC client
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "path/to/mailfort/pb" // Import your generated gRPC package

	"github.com/8thgencore/mailfort/internal/config"
)

func main() {
	// Create a gRPC connection to MailFort service
	conn, err := grpc.Dial("localhost:44044", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a MailFort client
	client := pb.NewMailFortClient(conn)

	// Example: Send Confirmation Email
	confirmationReq := &pb.ConfirmationRequest{
		Email: "user@example.com",
		Code:  "123456",
	}

	_, err = client.SendConfirmationEmail(context.Background(), confirmationReq)
	if err != nil {
		log.Fatalf("Failed to send confirmation email: %v", err)
	}

	// Example: Send Password Reset Email
	resetReq := &pb.ResetPasswordRequest{
		Email: "user@example.com",
		Code:  "654321",
	}

	_, err = client.SendPasswordResetEmail(context.Background(), resetReq)
	if err != nil {
		log.Fatalf("Failed to send password reset email: %v", err)
	}
}
```

### REST API Interaction
MailFort also provides a RESTful API for interaction. Below are examples using `curl`:

#### Send Confirmation Email:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com","code":"123456"}' http://localhost:8080/api/send-confirmation-email
```
#### Send Password Reset Email:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"email":"user@example.com","code":"654321"}' http://localhost:8080/api/send-password-reset-email
```

Make sure to replace the example email addresses and codes with your actual data.

Feel free to adapt the examples based on your programming language and preferred HTTP client library for REST API interaction.

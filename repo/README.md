# Goravel Application

This is a Goravel application generated from template.

## Getting Started

1. Copy `.env.example` to `.env` and configure your environment variables
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Run the application:
   ```bash
   go run main.go
   ```

## Project Structure

- `app/` - Application code
  - `http/` - HTTP controllers, middleware, resources
  - `console/` - Console commands
  - `grpc/` - gRPC controllers and interceptors
  - `jobs/` - Background jobs
  - `mails/` - Mail templates
  - `models/` - Database models
  - `providers/` - Service providers
  - `services/` - Business logic services
- `bootstrap/` - Application bootstrap
- `config/` - Configuration files
- `database/` - Migrations and seeders
- `routes/` - Route definitions
- `resources/` - Views and assets
- `storage/` - Logs and temporary files
- `tests/` - Test files

## Documentation

For more information, visit [Goravel Documentation](https://goravel.dev).


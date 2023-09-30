go build -o bookings.exe cmd/web/main.go cmd/web/routes.go cmd/web/middleware.go cmd/web/send-mail.go
bookings.exe -dbname=bookings -dbuser=postgres -dbpass=nimda -cache=false -production=false
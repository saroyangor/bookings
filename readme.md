# Bookings and Reservations

This is the repository for bookings and reservations project.

- Built in Go version 1.20.x

Dependencies:

- [chi router](https://github.com/go-chi/chi)
- [alex edwards SCS](https://github.com/alexedwards/scs) session management
- [nosurf](https://github.com/justinas/nosurf)
- [pgx](https://github.com/jackc/pgx)
- [simple mail](https://github.com/xhit/go-simple-mail)
- [Go validator](https://github.com/asaskevich/govalidator)

In order to build and run this application, it is necessary to
install Soda (go install github.com/gobuffalo/pop/... ), create
a postgres database, fill in the correct values in database.yml,
and then run soda migrate.

To build and run the application, from the root level of the project,
execute this command:
```
go build -o bookings ./cmd/web/ && ./bookings \
-dbname=bookings \
-dbuser=postgress
```
where you have the correct entires for your database name (dbName)
and database user (dbUser)
For the full list of command flags, run ./bookings -h

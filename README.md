# Muzz Backend Technical Exercise

A mini API that could power a very simple dating app.

## Assumptions

- I didn't need to support TLS/SSL.
- The DB ORM itself is checking for SQL injection.
- MariaDB was suitable as it was SQL.
- There was a discrepancy in the documentation for the 'user' object. Text asked for "date of birth" but the JSON example was "age". Age was simpler to implement so I went with that.
- Assumed that only the CRUD actions specified in the documentation were required. E.g. no editing of user details or changing swipe preferences.

## Notes on Data

The data is generated by https://github.com/brianvoe/gofakeit. This library is great for most things, but the location data is 'truly' random, whereby the latitude, longitude and name are generated independently. This means that the data is not geographically accurate. Please keep in mind that yes, it is very likely that the users you're looking at are 3000km away... or in the middle of the ocean. Although I agree it's not ideal, I think it's fine for the purposes of this exercise.

## Running the API

`docker compose up` will start the API and the database. The API will be available at `http://0.0.0.0:3000`. I've commented out the database ports, but they're default. You will need to hit `/user/create` a few times to generate your test data. I've included an `endpoints.json` file which was exported via Insomnia, but can be imported to Postman or similar. This export has some filters added by default for `/discovery` etc.

## Code Highlights / Notes

### The use of `init()`

I had used the `init()` function within `/user` to show how an alternative method could be used to call `AutoMigrate`. For code clarity I have also duplicated this function within `main.go` to match how the others are used.

### The `utils.Preserve` function

This is, what I think is, a very useful and extendable way to offer a graceful shutdown procedure. Due to the ORM used, I don't need to gracefully shut down the database as it will handle itself. But, if for example we wanted to notify an endpoint elsewhere we're shutting down, or want to ensure files are closed properly, this function could be extended to handle that.

### `/context`

This allows the ease of access to something like a database connection, a logger or prometheus client. It also IMO makes code much clearer as I don't have to pass the database instance around the app via structs or pointers within functions, I can simply request the context and get the database instance from it.

### Sequential IDs

By default I used the build-in functionality of SQL to auto-increment my IDs and matched the type specified in the documentation. However, in production I would use a UUID or similar non-sequential ID to avoid any potential security issues of a user trying to 'game' the system by guessing the next ID on the swipe endpoint.

---

I appreciate you taking the time to review my project!

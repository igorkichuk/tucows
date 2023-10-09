### Common for terminal and web execution:
- go to the directory of this file
- run `cp .env.template .env` to create `.env` file; change the values if you want to receive another size of the picture from the picture provider

#### To run from Terminal you have to:
 - make sure xterm256 is enabled in your terminal
 - run `go run ./client/cli -grayscale=1 -key=45` (you can skip the parameters)
 - change size and/or scale of the terminal and execute the program again to see the result of scaling

#### To see the result in Browser:
- run `go run ./client/web`
- open http://localhost:8080/post?grayscale=true&key=120 (query parameters are optional)

### Processing Errors:

- if a service provider is out of order you will see:
  -   errors description in terminal if you run the program in the terminal instead of image or quote
  - errors description in terminal and high level error description on the web page

### Running tests

- from the same directory run `go test ./...` (coverage is not full, but shows ability to work with different kinds of tests)

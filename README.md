<h1 align="center"> key-value style in-memory database. </h1>

<h2 align="center">a quick and easy to use k-v db built in go. </h2>

- includes cli support.
- persistant input/error handling for edge cases.
- interactive cli mode, which stores variable-like data (for now only supports values of type string and int) to then be retrieved later with methods like "GET", "SET", "DEL", and "EXIT" to exit the program.
- serializes values to []byte before storing to database struct for optimised performance.

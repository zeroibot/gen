# gen 

Tool for generating UUID, Code, and Hashed Passwords

## Usage 

`gen.exe (task) (params...)`

* `gen uuid` - generates a UUID v7 string
* `gen code (length) (pattern)` - generates a code of specified length; pattern uses **aA1** to indicate which characters to include
* `gen password (rawPassword)` - generates the hashed password using **bcrypt**
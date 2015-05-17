# Cassandra Auth

Simple demo app for demonstrate authorization and registration using Apache Cassandra.

## Running

1. Install docker:

    https://docs.docker.com/installation/

2. Install pip:

    https://pip.pypa.io/en/latest/installing.html

3. Install docker-compose:

        $ pip install -U docker-compose

4. Pull the app:

        $ git clone https://github.com/hypersleep/cassandra-auth.git && cd cassandra-auth

5. And run:

        $ docker-compose up

## Usage

`http` command in examples: https://github.com/jakubroztocil/httpie

Use `$(boot2docker ip)` on MacOS/Windows or `localhost` on Linux instead.

Register new user:

    $ http POST $(boot2docker ip):8080/users/create email=hypersleep@example.com password=qazxsw23
    HTTP/1.1 200 OK
    Content-Length: 61
    Content-Type: application/json
    Date: Sun, 17 May 2015 22:30:56 GMT

    {
        "data": {
            "message": "Successfully registred!"
        },
        "status": true
    }

Sign in:

    $ http POST $(boot2docker ip):8080/users/signin email=hypersleep@example.com password=qazxsw23
    HTTP/1.1 302 Found
    Content-Length: 61
    Content-Type: application/json
    Date: Sun, 17 May 2015 22:31:09 GMT
    Location: /users/check
    Set-Cookie: session=MTQzMTkwMTg2OXxEdi1CQkFFQ180SUFBUkFCRUFBQU5mLUNBQUVHYzNSeWFXNW5EQWNBQldWdFlXbHNCbk4wY21sdVp3d1lBQlpvZVhCbGNuTnNaV1Z3UUdWNFlXMXdiR1V1WTI5dHx1d7lPsKj8_BPspYUOKtHmJI8AHr1X_9BQuzTMe81aDQ==; Path=/; Expires=Tue, 16 Jun 2015 22:31:09 UTC; Max-Age=2592000

    {
        "data": {
            "message": "Successfully signed in!"
        },
        "status": true
    }

Copy session cookie from last response, paste as header in next command (without space between ':'' and 'session') and try to check your auth:

    $ http POST $(boot2docker ip):8080/users/check Cookie:session=MTQzMTkwMTg2OXxEdi1CQkFFQ180SUFBUkFCRUFBQU5mLUNBQUVHYzNSeWFXNW5EQWNBQldWdFlXbHNCbk4wY21sdVp3d1lBQlpvZVhCbGNuTnNaV1Z3UUdWNFlXMXdiR1V1WTI5dHx1d7lPsKj8_BPspYUOKtHmJI8AHr1X_9BQuzTMe81aDQ==;
    HTTP/1.1 200 OK
    Content-Length: 90
    Content-Type: application/json
    Date: Sun, 17 May 2015 22:31:42 GMT

    {
        "data": {
            "message": "Successfully checked! Hello, hypersleep@example.com!"
        },
        "status": true
    }

Try to log out and copy changed session in clipboard again:

    $ http DELETE $(boot2docker ip):8080/users/logout Cookie:session=MTQzMTkwMTg2OXxEdi1CQkFFQ180SUFBUkFCRUFBQU5mLUNBQUVHYzNSeWFXNW5EQWNBQldWdFlXbHNCbk4wY21sdVp3d1lBQlpvZVhCbGNuTnNaV1Z3UUdWNFlXMXdiR1V1WTI5dHx1d7lPsKj8_BPspYUOKtHmJI8AHr1X_9BQuzTMe81aDQ==;
    HTTP/1.1 200 OK
    Content-Length: 62
    Content-Type: application/json
    Date: Sun, 17 May 2015 22:32:10 GMT
    Set-Cookie: session=MTQzMTkwMTkzMHxEdi1CQkFFQ180SUFBUkFCRUFBQUZmLUNBQUVHYzNSeWFXNW5EQWNBQldWdFlXbHNBQT09fGG1n_6ckxLF5AJ3JVaktebJEtjKHIu5GOzgA-v0HO7d; Path=/; Expires=Tue, 16 Jun 2015 22:32:10 UTC; Max-Age=2592000

    {
        "data": {
            "message": "Successfully logged out!"
        },
        "status": true
    }

Check new session again, and you'll got an error:

    $ http POST $(boot2docker ip):8080/users/check Cookie:session=MTQzMTkwMTkzMHxEdi1CQkFFQ180SUFBUkFCRUFBQUZmLUNBQUVHYzNSeWFXNW5EQWNBQldWdFlXbHNBQT09fGG1n_6ckxLF5AJ3JVaktebJEtjKHIu5GOzgA-v0HO7d;
    HTTP/1.1 200 OK
    Content-Length: 44
    Content-Type: application/json
    Date: Sun, 17 May 2015 22:32:41 GMT

    {
        "data": {
            "message": "Error!"
        },
        "status": true
    }
